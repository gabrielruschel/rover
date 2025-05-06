package navigation

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"
)

func parsePlateau(coordStr string) (upperX uint64, upperY uint64, err error) {
	splitStr := strings.Split(coordStr, " ")
	if len(splitStr) < 2 {
		err = fmt.Errorf("could not process plateau upper coordinates")
		return
	}

	upperX, err = strconv.ParseUint(splitStr[0], 10, 64)
	if err != nil {
		err = fmt.Errorf("could not parse upper X coordinate [%s]: %w", splitStr[0], err)
		return
	}

	upperY, err = strconv.ParseUint(splitStr[1], 10, 64)
	if err != nil {
		err = fmt.Errorf("could not parse upper Y coordinate [%s]: %w", splitStr[1], err)
	}

	return
}

func NavigateRovers(input io.Reader, logger *slog.Logger) (output []string, err error) {
	fileScanner := bufio.NewScanner(input)

	if !fileScanner.Scan() {
		err = fmt.Errorf("could not read first line of input file: %v", fileScanner.Err())
		return
	}

	platLine := fileScanner.Text()
	upperX, upperY, err := parsePlateau(platLine)
	if err != nil {
		err = fmt.Errorf("error while parsing plateau upper coordinates: %w", err)
		return
	}
	logger.Debug("parsed upper-right plateau coordinates", slog.Uint64("upperX", upperX), slog.Uint64("upperY", upperY))

	return
}
