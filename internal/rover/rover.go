package rover

import (
	"fmt"
	"log/slog"
	"strings"
	"unicode"

	"github.com/gabrielruschel/rover/internal/helpers"
)

func parseRoverPosition(posStr string) (
	coordX uint64, coordY uint64,
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

	orientRunes := []rune(split[2])
	if len(orientRunes) == 0 {
		err = fmt.Errorf("could not parse rover position orientation")
		return
	}
	orientation = unicode.ToUpper(orientRunes[0])

	switch orientation {
	case 'N':
	case 'S':
	case 'E':
	case 'W':
	default:
		err = fmt.Errorf("invalid orientation [%c]", orientation)
	}

	return
}

type Rover struct {
	XCoord      uint64
	YCoord      uint64
	Orientation rune

	logger *slog.Logger
}

func NewRover(posStr string, logger *slog.Logger) (*Rover, error) {
	coordX, coordY, orientation, err := parseRoverPosition(posStr)
	if err != nil {
		return nil, err
	}

	logger.Info("created rover", slog.Uint64("X", coordX), slog.Uint64("Y", coordY), slog.Any("orientation", orientation))

	return &Rover{
		XCoord:      coordX,
		YCoord:      coordY,
		Orientation: orientation,
		logger:      logger,
	}, nil
}

func (r Rover) ExecuteRoverNavigation(instStr string) {
}
