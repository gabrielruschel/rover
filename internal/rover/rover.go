package rover

import (
	"fmt"
	"log/slog"
	"strings"
	"unicode"

	"github.com/gabrielruschel/rover/internal/helpers"
)

func parseRoverPosition(
	posStr string, upperX, upperY uint64,
) (
	coordX, coordY uint64,
	orientation rune, err error,
) {
	split := strings.Split(posStr, " ")
	if len(split) != 3 {
		err = fmt.Errorf("could not parse rover position, not enough info")
		return
	}

	coordX, coordY, err = helpers.ParseCoordinates(strings.Join(split[:2], " "))
	if err != nil {
		err = fmt.Errorf("could not parse rover position: %w", err)
		return
	}

	if coordX < 0 || coordY < 0 || coordX > upperX || coordY > upperY {
		err = fmt.Errorf("cannot place rover in position (%d,%d): out of bounds", coordX, coordY)
		return
	}

	orientRunes := []rune(split[2])
	if len(orientRunes) == 0 {
		err = fmt.Errorf("could not parse rover position orientation")
		return
	}
	orientation = unicode.ToUpper(orientRunes[0])

	switch orientation {
	case North:
	case South:
	case East:
	case West:
	default:
		err = fmt.Errorf("invalid orientation [%c]", orientation)
	}

	return
}

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
	logger *slog.Logger,
) (*Rover, error) {
	coordX, coordY, orientation, err := parseRoverPosition(posStr, upperX, upperY)
	if err != nil {
		return nil, err
	}

	logger.Info("created rover", slog.Uint64("X", coordX), slog.Uint64("Y", coordY), slog.Any("orientation", orientation))

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
		r.logger.Debug("executing instruction", slog.Any("instruction", inst))

		switch inst {
		case Left, Right:
			r.rotateRover(inst)
			r.logger.Info("rotated rover", slog.Any("orientation", r.Orientation))
		case 'M':
			err := r.moveRover(deployedRovers)
			if err != nil {
				r.logger.Warn(err.Error())
				continue
			}
			r.logger.Info("moved rover to position", slog.Uint64("X", r.XCoord), slog.Uint64("Y", r.YCoord))
		default:
			r.logger.Warn("invalid instruction, skipping", slog.Any("instruction", inst))
		}
	}

	fmt.Printf("%d %d %c\n", r.XCoord, r.YCoord, r.Orientation)
	return r.XCoord, r.YCoord
}

func (r *Rover) rotateRover(direction rune) {
	switch direction {
	case Left:
		r.orientationIdx--
		if r.orientationIdx < 0 {
			r.orientationIdx = len(Orientations) - 1
		}
	case Right:
		r.orientationIdx++
		if r.orientationIdx > len(Orientations)-1 {
			r.orientationIdx = 0
		}
	}
	r.Orientation = Orientations[r.orientationIdx]
}

func (r *Rover) moveRover(deployedRovers [][2]uint64) (err error) {
	newX, newY := r.XCoord, r.YCoord
	switch r.Orientation {
	case North:
		newY++
	case South:
		newY--
	case East:
		newX++
	case West:
		newX--
	}

	err = r.validateDestination(newX, newY, deployedRovers)
	if err != nil {
		return
	}

	r.XCoord, r.YCoord = newX, newY
	return
}

func (r Rover) validateDestination(
	newX, newY uint64,
	deployedRovers [][2]uint64,
) (err error) {
	if newX < 0 || newY < 0 || newX > r.upperX || newY > r.upperY {
		err = fmt.Errorf("cannot move rover to (%d,%d): invalid position, out of bounds", newX, newY)
		return
	}

	for i, dr := range deployedRovers {
		if newX == dr[0] && newY == dr[1] {
			err = fmt.Errorf("cannot move rover to (%d,%d): invalid position, already occupied by rover r%d", newX, newY, i)
			break
		}
	}

	return
}
