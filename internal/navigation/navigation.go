package navigation

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"

	"github.com/gabrielruschel/rover/internal/helpers"
)

func NavigateRovers(input io.Reader, logger *slog.Logger) (output []string, err error) {
	fileScanner := bufio.NewScanner(input)

	if !fileScanner.Scan() {
		err = fmt.Errorf("could not read first line of input file: %v", fileScanner.Err())
		return
	}

	platLine := fileScanner.Text()
	upperX, upperY, err := helpers.ParseCoordinates(platLine)
	if err != nil {
		err = fmt.Errorf("error while parsing plateau upper coordinates: %w", err)
		return
	}
	logger.Debug("parsed upper-right plateau coordinates", slog.Uint64("upperX", upperX), slog.Uint64("upperY", upperY))

	return
}
