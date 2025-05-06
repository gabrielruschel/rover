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
		logger.Error(fmt.Sprintf("unexpected error opening input file %s: %v", runConfig.InputFile, err))
		return
	}
	defer inputFile.Close()

	logger.Debug("starting navigation")
	output, err := navigation.NavigateRovers(inputFile, logger)
}
