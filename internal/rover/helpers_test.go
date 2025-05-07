package rover

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRoverPosition(t *testing.T) {
	tCases := []struct {
		name        string
		posStr      string
		coordX      uint64
		coordY      uint64
		orientation rune
		err         bool
	}{
		{
			name:        "Test valid position",
			posStr:      "5 7 N",
			coordX:      5,
			coordY:      7,
			orientation: 'N',
		},
		{
			name:        "Test valid position with lowercase orientation",
			posStr:      "99 5000 w",
			coordX:      99,
			coordY:      5000,
			orientation: 'W',
		},
		{
			name:   "Test invalid position string",
			posStr: "S 33 44",
			err:    true,
		},
		{
			name:   "Test invalid position with negative coordinate",
			posStr: "-1 6 w",
			err:    true,
		},
		{
			name:   "Test invalid position with invalid orientation",
			posStr: "2 59 O",
			err:    true,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			resX, resY, resOrient, err := parseRoverPosition(tc.posStr)
			if tc.err {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tc.coordX, resX)
			assert.Equal(t, tc.coordY, resY)
			assert.Equal(t, tc.orientation, resOrient)
		})
	}
}

func TestValidatePosition(t *testing.T) {
	tCases := []struct {
		name           string
		destX          uint64
		destY          uint64
		upperX         uint64
		upperY         uint64
		deployedRovers [][2]uint64
		err            error
	}{
		{
			name:   "Test valid position",
			destX:  5,
			destY:  7,
			upperX: 10,
			upperY: 11,
			deployedRovers: [][2]uint64{
				{5, 3},
			},
		},
		{
			name:   "Test invalid position X out of bounds",
			destX:  56,
			destY:  90,
			upperX: 42,
			upperY: 99,
			err:    fmt.Errorf("cannot move rover to (56,90): invalid position, out of bounds"),
		},
		{
			name:   "Test invalid position Y out of bounds",
			destX:  50,
			destY:  96,
			upperX: 57,
			upperY: 25,
			err:    fmt.Errorf("cannot move rover to (50,96): invalid position, out of bounds"),
		},
		{
			name:   "Test invalid position already occupied",
			destX:  23,
			destY:  23,
			upperX: 57,
			upperY: 25,
			deployedRovers: [][2]uint64{
				{2, 1},
				{23, 23},
			},
			err: fmt.Errorf("cannot move rover to (23,23): invalid position, already occupied by rover r1"),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateDestination(tc.destX, tc.destY, tc.upperX, tc.upperY, tc.deployedRovers)
			assert.Equal(t, tc.err, err)
		})
	}
}
