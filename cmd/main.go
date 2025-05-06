package main

import (
	"fmt"

	"github.com/gabrielruschel/rover/internal/config"
)

func main() {
	fmt.Println("starting rover!")

	runConfig := config.NewConfig()
	fmt.Printf("loaded config = %+v\n", runConfig)
}
