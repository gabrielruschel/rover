package navigation

import (
	"strings"
	"testing"

	"github.com/gabrielruschel/rover/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestNavigateRovers(t *testing.T) {
	tCases := []struct {
		name      string
		fileInput string
		output    []string
		err       bool
	}{
		{
			name:      "Test valid navigation",
			fileInput: validNavigation1,
			output:    []string{"1 3 N", "5 1 E"},
		},
		{
			name:      "Test valid navigation two rovers",
			fileInput: validNavigation2,
			output:    []string{"5 20 N", "15 0 S"},
		},
		{
			name:      "Test valid navigation with rover deploy fail",
			fileInput: validNavigationWithRoverFailInvalidPosition,
			output:    []string{"1 3 N", "2 0 W"},
		},
		{
			name:      "Test invalid navigation all rovers failing to deploy",
			fileInput: invalidNavigationAllRoversFailingToDeploy,
			output:    nil,
		},
		{
			name:      "Test valid navigation with rover deploy fail out of bounds",
			fileInput: validNavigationWithRoverFailInvalidCoordinatesOutOfBounds,
			output:    []string{"1 3 N"},
		},
		{
			name:      "Test valid navigation with rover deploy fail negative coordinates",
			fileInput: validNavigationWithRoverFailInvalidCoordinatesNegative,
			output:    []string{"1 3 N"},
		},
		{
			name:      "Test valid navigation with rover deploy fail occupied coordinates",
			fileInput: validNavigationWithRoverFailInvalidOccupiedCoordinates,
			output:    []string{"1 3 N"},
		},
		{
			name:      "Test valid navigation with rover going out of bounds",
			fileInput: validNavigationWithRoverGoingOutOfBounds,
			output:    []string{"1 3 N", "5 3 E", "2 0 W", "4 5 N"},
		},
		{
			name:      "Test valid navigation with rover going out of bounds negative coordinates",
			fileInput: validNavigationWithRoverGoingOutOfBoundsNegative,
			output:    []string{"1 3 N", "0 2 W"},
		},
		{
			name:      "Test valid navigation with rover going out of bounds upper overflow",
			fileInput: validNavigationWithRoverGoingOutOfBoundsUpperOverflow,
			output:    []string{"1 3 N", "1 18446744073709551615 N"},
		},
		{
			name:      "Test valid navigation with rover hitting occupied position",
			fileInput: validNavigationWithRoverHittingOccupiedPosition,
			output:    []string{"1 3 N", "2 0 W", "1 2 N"},
		},
		{
			name:      "Test valid navigation with invalid instructions",
			fileInput: validNavigationWithInvalidInstructions,
			output:    []string{"1 3 N", "5 1 E"},
		},
		{
			name:      "Test valid navigation with missing instructions",
			fileInput: validNavigationWithMissingInstructions,
			output:    []string{"1 3 N", "3 3 E"},
		},
		{
			name:      "Test invalid plateau input",
			fileInput: invalidPlateauCoordinates,
			err:       true,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.fileInput)
			logger := helpers.NewLogger("DEBUG")

			res, err := NavigateRovers(reader, logger)
			if tc.err {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tc.output, res)
		})
	}
}
