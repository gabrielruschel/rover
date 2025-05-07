package rover

import (
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
			name:   "Test invalid position",
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

func TestRotateRover(t *testing.T) {
	tCases := []struct {
		name              string
		direction         rune
		currOrient        rune
		currOrientIdx     int
		expectedOrient    rune
		expectedOrientIdx int
	}{
		{
			name:              "Test rotate LEFT from NORTH",
			direction:         Left,
			currOrient:        North,
			currOrientIdx:     OrientationIdx[North],
			expectedOrient:    West,
			expectedOrientIdx: OrientationIdx[West],
		},
		{
			name:              "Test rotate LEFT from WEST",
			direction:         Left,
			currOrient:        West,
			currOrientIdx:     OrientationIdx[West],
			expectedOrient:    South,
			expectedOrientIdx: OrientationIdx[South],
		},
		{
			name:              "Test rotate LEFT from SOUTH",
			direction:         Left,
			currOrient:        South,
			currOrientIdx:     OrientationIdx[South],
			expectedOrient:    East,
			expectedOrientIdx: OrientationIdx[East],
		},
		{
			name:              "Test rotate LEFT from EAST",
			direction:         Left,
			currOrient:        East,
			currOrientIdx:     OrientationIdx[East],
			expectedOrient:    North,
			expectedOrientIdx: OrientationIdx[North],
		},

		// ---------------------------------------------------------

		{
			name:              "Test rotate RIGHT from NORTH",
			direction:         Right,
			currOrient:        North,
			currOrientIdx:     OrientationIdx[North],
			expectedOrient:    East,
			expectedOrientIdx: OrientationIdx[East],
		},
		{
			name:              "Test rotate RIGHT from EAST",
			direction:         Right,
			currOrient:        East,
			currOrientIdx:     OrientationIdx[East],
			expectedOrient:    South,
			expectedOrientIdx: OrientationIdx[South],
		},
		{
			name:              "Test rotate RIGHT from SOUTH",
			direction:         Right,
			currOrient:        South,
			currOrientIdx:     OrientationIdx[South],
			expectedOrient:    West,
			expectedOrientIdx: OrientationIdx[West],
		},
		{
			name:              "Test rotate RIGHT from WEST",
			direction:         Right,
			currOrient:        West,
			currOrientIdx:     OrientationIdx[West],
			expectedOrient:    North,
			expectedOrientIdx: OrientationIdx[North],
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			rov := &Rover{
				Orientation:    tc.currOrient,
				orientationIdx: tc.currOrientIdx,
			}

			rov.rotateRover(tc.direction)
			assert.Equal(t, tc.expectedOrient, rov.Orientation)
			assert.Equal(t, tc.expectedOrientIdx, rov.orientationIdx)
		})
	}
}
