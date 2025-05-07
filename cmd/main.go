package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gabrielruschel/rover/internal/config"
	"github.com/gabrielruschel/rover/internal/helpers"
	"github.com/gabrielruschel/rover/internal/navigation"
)

func main() {
	runConfig := config.NewConfig()
	logger := helpers.NewLogger(runConfig.LogLevel, slog.String("job", "navigator"))
	logger.Debug("loaded config", slog.Any("config", runConfig))

	inputFile, err := os.Open(runConfig.InputFile)
	if err != nil {
		logger.Error(fmt.Sprintf("unexpected error opening input file %s: %s", runConfig.InputFile, err.Error()))
		return
	}
	defer inputFile.Close()

	logger.Debug("starting navigation")
	roverPos, err := navigation.NavigateRovers(inputFile, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("an error occurred while processing rover navigation: %s", err.Error()))
		return
	}

	// also write rover navigation results to output file
	outFile, err := os.OpenFile(runConfig.OutputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o744)
	defer outFile.Close()

	for _, rov := range roverPos {
		if _, errWrite := fmt.Fprintln(outFile, rov); errWrite != nil {
			logger.Error(fmt.Sprintf("unexpected error writing result file: %s", err.Error()))
			return
		}
	}
}
