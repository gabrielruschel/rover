package navigation

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"

	"github.com/gabrielruschel/rover/internal/helpers"
	"github.com/gabrielruschel/rover/internal/rover"
)

func NavigateRovers(input io.Reader, logger *slog.Logger) (output []string, err error) {
	fileScanner := bufio.NewScanner(input)

	// read file first line (plateau upper coordinates)
	if !fileScanner.Scan() {
		err = fmt.Errorf("could not read first line of input file: %v", fileScanner.Err())
		return
	}

	platLine := fileScanner.Text()
	upperX, upperY, err := helpers.ParseUint64Coordinates(platLine)
	if err != nil {
		err = fmt.Errorf("error while parsing plateau upper coordinates: %w", err)
		return
	}
	logger.Debug("parsed upper-right plateau coordinates", slog.Uint64("upperX", upperX), slog.Uint64("upperY", upperY))

	var (
		counter        uint64
		deployedRovers = make([][2]uint64, 0, int(upperX))
	)

	for {
		// read first line (initial rover position)
		read := fileScanner.Scan()
		if !read {
			err = helpers.CheckScannerError(fileScanner)
			if err != nil {
				return
			}

			// reached end of file
			break
		}

		posLine := fileScanner.Text()
		rov, errRov := rover.NewRover(
			posLine,
			upperX, upperY,
			deployedRovers,
			logger.With(slog.String("rover", fmt.Sprintf("r%d", counter))),
		)
		if errRov != nil {
			logger.Error(
				"error while creating rover, skipping",
				slog.String("err", errRov.Error()),
				slog.String("rover", fmt.Sprintf("r%d", counter)),
			)
			fileScanner.Scan() // read next line, ignore instructions for this failed rover
			continue
		}

		// read second line (rover navigation instructions)
		read = fileScanner.Scan()
		if !read {
			err = helpers.CheckScannerError(fileScanner)
			if err != nil {
				return
			}
			logger.Warn("did not receive instructions for rover")
		}
		instLine := fileScanner.Text()

		finalPosX, finalPosY := rov.ExecuteRoverNavigation(instLine, deployedRovers)
		deployedRovers = append(deployedRovers, [2]uint64{finalPosX, finalPosY})
		output = append(output, fmt.Sprintf("%d %d %c", finalPosX, finalPosY, rov.Orientation))

		counter++
	}

	return
}
