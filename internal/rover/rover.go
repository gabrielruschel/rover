package rover

import (
	"fmt"
	"log/slog"
	"unicode"
)

type Rover struct {
	XCoord         uint64
	YCoord         uint64
	Orientation    rune
	orientationIdx int

	upperX uint64
	upperY uint64

	logger *slog.Logger
}

func NewRover(
	posStr string,
	upperX, upperY uint64,
	deployedRovers [][2]uint64,
	logger *slog.Logger,
) (*Rover, error) {
	coordX, coordY, orientation, err := parseRoverPosition(posStr)
	if err != nil {
		return nil, err
	}

	err = validateDestination(coordX, coordY, upperX, upperY, deployedRovers)
	if err != nil {
		return nil, err
	}

	logger.Info(
		"created rover",
		slog.Uint64("X", coordX), slog.Uint64("Y", coordY),
		slog.String("orientation", fmt.Sprintf("%c", orientation)),
	)

	return &Rover{
		XCoord:         coordX,
		YCoord:         coordY,
		Orientation:    orientation,
		orientationIdx: OrientationIdx[orientation],
		upperX:         upperX,
		upperY:         upperY,
		logger:         logger,
	}, nil
}

func (r *Rover) ExecuteRoverNavigation(
	instStr string,
	deployedRovers [][2]uint64,
) (uint64, uint64) {
	for _, inst := range instStr {
		inst = unicode.ToUpper(inst)
		r.logger.Debug("executing instruction", slog.String("instruction", fmt.Sprintf("%c", inst)))

		switch inst {
		case RotateLeft, RotateRight:
			r.rotateRover(inst)
			r.logger.Info("rotated rover", slog.String("orientation", fmt.Sprintf("%c", r.Orientation)))
		case Move:
			err := r.moveRover(deployedRovers)
			if err != nil {
				r.logger.Warn(err.Error())
				continue
			}
			r.logger.Info(
				"moved rover to position",
				slog.Uint64("X", r.XCoord), slog.Uint64("Y", r.YCoord),
				slog.String("orientation", fmt.Sprintf("%c", r.Orientation)),
			)
		default:
			r.logger.Warn("invalid instruction, skipping", slog.String("instruction", fmt.Sprintf("%c", inst)))
		}
	}

	return r.XCoord, r.YCoord
}

func (r *Rover) rotateRover(direction rune) {
	switch direction {
	case RotateLeft:
		r.orientationIdx--
		if r.orientationIdx < 0 {
			r.orientationIdx = len(Orientations) - 1
		}
	case RotateRight:
		r.orientationIdx++
		if r.orientationIdx > len(Orientations)-1 {
			r.orientationIdx = 0
		}
	}
	r.Orientation = Orientations[r.orientationIdx]
}

func (r *Rover) moveRover(deployedRovers [][2]uint64) (err error) {
	newX, newY := r.XCoord, r.YCoord

	// treating each corner case individually to avoid problems with uint64 overflow
	switch r.Orientation {
	case North:
		newY++
		// uint64 overflow comes back to 0
		if newY == 0 {
			err = fmt.Errorf("cannot move rover to (%d,%d): invalid position, out of bounds", newX, newY)
			return
		}
	case South:
		if r.YCoord == 0 {
			err = fmt.Errorf("cannot move rover to (%d,-1): invalid position, out of bounds", newX)
			return
		}
		newY--
	case East:
		newX++
		// uint64 overflow comes back to 0
		if newX == 0 {
			err = fmt.Errorf("cannot move rover to (%d,%d): invalid position, out of bounds", newX, newY)
			return
		}
	case West:
		if r.XCoord == 0 {
			err = fmt.Errorf("cannot move rover to (-1,%d): invalid position, out of bounds", newY)
			return
		}
		newX--
	}

	err = validateDestination(newX, newY, r.upperX, r.upperY, deployedRovers)
	if err != nil {
		return
	}

	r.XCoord, r.YCoord = newX, newY
	return
}
