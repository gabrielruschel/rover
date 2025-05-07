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
