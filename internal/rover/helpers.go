package rover

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/gabrielruschel/rover/internal/helpers"
)

const (
	North = 'N'
	South = 'S'
	East  = 'E'
	West  = 'W'

	RotateLeft  = 'L'
	RotateRight = 'R'
	Move        = 'M'
)

var (
	Orientations   = [4]rune{East, South, West, North}
	OrientationIdx = map[rune]int{
		East:  0,
		South: 1,
		West:  2,
		North: 3,
	}
)

func parseRoverPosition(posStr string) (
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

func validateDestination(
	newX, newY,
	upperX, upperY uint64,
	deployedRovers [][2]uint64,
) (err error) {
	if newX > upperX || newY > upperY {
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
