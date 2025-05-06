package helpers

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func NewLogger(level string, attrs ...any) *slog.Logger {
	opts := &slog.HandlerOptions{}
	switch strings.ToUpper(level) {
	case "DEBUG":
		opts.Level = slog.LevelDebug
	case "INFO":
		opts.Level = slog.LevelInfo
	case "ERROR":
		opts.Level = slog.LevelError
	default:
		opts.Level = slog.LevelWarn
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, opts)).With(attrs...)
}

func ParseCoordinates(coordStr string) (upperX uint64, upperY uint64, err error) {
	splitStr := strings.Split(coordStr, " ")
	if len(splitStr) < 2 {
		err = fmt.Errorf("could not parse coordinates, not enough info")
		return
	}

	upperX, err = strconv.ParseUint(splitStr[0], 10, 64)
	if err != nil {
		err = fmt.Errorf("could not parse X coordinate [%s]: %w", splitStr[0], err)
		return
	}

	upperY, err = strconv.ParseUint(splitStr[1], 10, 64)
	if err != nil {
		err = fmt.Errorf("could not parse Y coordinate [%s]: %w", splitStr[1], err)
	}

	return
}

func CheckScannerError(scanner *bufio.Scanner) (err error) {
	if scanErr := scanner.Err(); scanErr != nil {
		err = fmt.Errorf("unexpected error reading line: %w", scanErr)
	}
	return
}
