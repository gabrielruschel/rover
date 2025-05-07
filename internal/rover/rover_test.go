package rover

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			direction:         RotateLeft,
			currOrient:        North,
			currOrientIdx:     OrientationIdx[North],
			expectedOrient:    West,
			expectedOrientIdx: OrientationIdx[West],
		},
		{
			name:              "Test rotate LEFT from WEST",
			direction:         RotateLeft,
			currOrient:        West,
			currOrientIdx:     OrientationIdx[West],
			expectedOrient:    South,
			expectedOrientIdx: OrientationIdx[South],
		},
		{
			name:              "Test rotate LEFT from SOUTH",
			direction:         RotateLeft,
			currOrient:        South,
			currOrientIdx:     OrientationIdx[South],
			expectedOrient:    East,
			expectedOrientIdx: OrientationIdx[East],
		},
		{
			name:              "Test rotate LEFT from EAST",
			direction:         RotateLeft,
			currOrient:        East,
			currOrientIdx:     OrientationIdx[East],
			expectedOrient:    North,
			expectedOrientIdx: OrientationIdx[North],
		},

		// ---------------------------------------------------------

		{
			name:              "Test rotate RIGHT from NORTH",
			direction:         RotateRight,
			currOrient:        North,
			currOrientIdx:     OrientationIdx[North],
			expectedOrient:    East,
			expectedOrientIdx: OrientationIdx[East],
		},
		{
			name:              "Test rotate RIGHT from EAST",
			direction:         RotateRight,
			currOrient:        East,
			currOrientIdx:     OrientationIdx[East],
			expectedOrient:    South,
			expectedOrientIdx: OrientationIdx[South],
		},
		{
			name:              "Test rotate RIGHT from SOUTH",
			direction:         RotateRight,
			currOrient:        South,
			currOrientIdx:     OrientationIdx[South],
			expectedOrient:    West,
			expectedOrientIdx: OrientationIdx[West],
		},
		{
			name:              "Test rotate RIGHT from WEST",
			direction:         RotateRight,
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

func TestMoveRover(t *testing.T) {
	tCases := []struct {
		name           string
		direction      rune
		coordX         uint64
		coordY         uint64
		upperX         uint64
		upperY         uint64
		deployedRovers [][2]uint64
		destX          uint64
		destY          uint64
		err            error
	}{
		{
			name:           "Test move rover NORTH",
			direction:      North,
			coordX:         2,
			coordY:         3,
			upperX:         50,
			upperY:         50,
			destX:          2,
			destY:          4,
			deployedRovers: [][2]uint64{{9, 9}},
		},
		{
			name:           "Test move rover SOUTH",
			direction:      South,
			coordX:         2,
			coordY:         3,
			upperX:         50,
			upperY:         50,
			destX:          2,
			destY:          2,
			deployedRovers: [][2]uint64{{9, 9}},
		},
		{
			name:           "Test move rover EAST",
			direction:      East,
			coordX:         2,
			coordY:         3,
			upperX:         50,
			upperY:         50,
			destX:          3,
			destY:          3,
			deployedRovers: [][2]uint64{{9, 9}},
		},
		{
			name:           "Test move rover WEST",
			direction:      West,
			coordX:         2,
			coordY:         3,
			upperX:         50,
			upperY:         50,
			destX:          1,
			destY:          3,
			deployedRovers: [][2]uint64{{9, 9}},
		},
		{
			name:           "Test move rover NORTH error out of bounds",
			direction:      North,
			coordX:         2,
			coordY:         50,
			upperX:         50,
			upperY:         50,
			destX:          2,
			destY:          50,
			deployedRovers: [][2]uint64{{9, 9}},
			err:            fmt.Errorf("cannot move rover to (2,51): invalid position, out of bounds"),
		},
		{
			name:           "Test move rover NORTH error out of bounds overflow",
			direction:      North,
			coordX:         2,
			coordY:         18446744073709551615,
			upperX:         50,
			upperY:         18446744073709551615,
			destX:          2,
			destY:          18446744073709551615,
			deployedRovers: [][2]uint64{{9, 9}},
			err:            fmt.Errorf("cannot move rover to (2,0): invalid position, out of bounds"),
		},
		{
			name:           "Test move rover SOUTH error out of bounds",
			direction:      South,
			coordX:         2,
			coordY:         0,
			upperX:         50,
			upperY:         50,
			destX:          2,
			destY:          0,
			deployedRovers: [][2]uint64{{9, 9}},
			err:            fmt.Errorf("cannot move rover to (2,-1): invalid position, out of bounds"),
		},
		{
			name:           "Test move rover EAST error out of bounds",
			direction:      East,
			coordX:         50,
			coordY:         3,
			upperX:         50,
			upperY:         50,
			destX:          50,
			destY:          3,
			deployedRovers: [][2]uint64{{9, 9}},
			err:            fmt.Errorf("cannot move rover to (51,3): invalid position, out of bounds"),
		},
		{
			name:           "Test move rover EAST error out of bounds overflow",
			direction:      East,
			coordX:         18446744073709551615,
			coordY:         3,
			upperX:         18446744073709551615,
			upperY:         50,
			destX:          18446744073709551615,
			destY:          3,
			deployedRovers: [][2]uint64{{9, 9}},
			err:            fmt.Errorf("cannot move rover to (0,3): invalid position, out of bounds"),
		},
		{
			name:           "Test move rover WEST error out of bounds",
			direction:      West,
			coordX:         0,
			coordY:         3,
			upperX:         50,
			upperY:         50,
			destX:          0,
			destY:          3,
			deployedRovers: [][2]uint64{{9, 9}},
			err:            fmt.Errorf("cannot move rover to (-1,3): invalid position, out of bounds"),
		},
		{
			name:           "Test move rover NORTH error position occupied",
			direction:      North,
			coordX:         2,
			coordY:         3,
			upperX:         50,
			upperY:         50,
			destX:          2,
			destY:          3,
			deployedRovers: [][2]uint64{{9, 9}, {2, 4}},
			err:            fmt.Errorf("cannot move rover to (2,4): invalid position, already occupied by rover r1"),
		},
		{
			name:           "Test move rover SOUTH error position occupied",
			direction:      South,
			coordX:         2,
			coordY:         3,
			upperX:         50,
			upperY:         50,
			destX:          2,
			destY:          3,
			deployedRovers: [][2]uint64{{9, 9}, {2, 2}},
			err:            fmt.Errorf("cannot move rover to (2,2): invalid position, already occupied by rover r1"),
		},
		{
			name:           "Test move rover EAST error position occupied",
			direction:      East,
			coordX:         2,
			coordY:         3,
			upperX:         50,
			upperY:         50,
			destX:          2,
			destY:          3,
			deployedRovers: [][2]uint64{{9, 9}, {3, 3}},
			err:            fmt.Errorf("cannot move rover to (3,3): invalid position, already occupied by rover r1"),
		},
		{
			name:           "Test move rover WEST error position occupied",
			direction:      West,
			coordX:         2,
			coordY:         3,
			upperX:         50,
			upperY:         50,
			destX:          2,
			destY:          3,
			deployedRovers: [][2]uint64{{9, 9}, {1, 3}},
			err:            fmt.Errorf("cannot move rover to (1,3): invalid position, already occupied by rover r1"),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			rov := &Rover{
				XCoord:      tc.coordX,
				YCoord:      tc.coordY,
				upperX:      tc.upperX,
				upperY:      tc.upperY,
				Orientation: tc.direction,
			}

			err := rov.moveRover(tc.deployedRovers)
			assert.Equal(t, tc.destX, rov.XCoord)
			assert.Equal(t, tc.destY, rov.YCoord)
			assert.Equal(t, tc.err, err)
		})
	}
}
