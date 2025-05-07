package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePlateau(t *testing.T) {
	testCases := []struct {
		name     string
		coordStr string
		upperX   uint64
		upperY   uint64
		err      bool
	}{
		{
			name:     "Test valid coordinates",
			coordStr: "5 7",
			upperX:   5,
			upperY:   7,
		},
		{
			name:     "Test valid coordinates with garbage after y",
			coordStr: "5 5 76596896789",
			upperX:   5,
			upperY:   5,
		},
		{
			name:     "Test invalid coordinates",
			coordStr: "55",
			err:      true,
		},
		{
			name:     "Test invalid negative coordinates",
			coordStr: "-2 -3",
			err:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uppX, uppY, err := ParseUint64Coordinates(tc.coordStr)
			if tc.err {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tc.upperX, uppX)
			assert.Equal(t, tc.upperY, uppY)
		})
	}
}
