package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gabrielruschel/rover/internal/config"
	"github.com/gabrielruschel/rover/internal/navigation"
)

func main() {
	fmt.Println("starting rover!")

	runConfig := config.NewConfig()
	fmt.Printf("loaded config = %+v\n", runConfig)

	inputFile, err := os.Open(runConfig.InputFile)
	if err != nil {
		log.Fatalf("unexpected error opening input file %s: %v", runConfig.InputFile, err)
	}
	defer inputFile.Close()

	output, err := navigation.NavigateRovers(inputFile)
}
