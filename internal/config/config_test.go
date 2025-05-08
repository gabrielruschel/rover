package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	testCases := []struct {
		name       string
		inputFile  string
		outputFile string
		logLevel   string
		expected   Config
	}{
		{
			name:       "Test default config",
			inputFile:  "",
			outputFile: "",
			expected: Config{
				InputFile:  "input.txt",
				OutputFile: "output.txt",
				LogLevel:   "ERROR",
			},
		},
		{
			name:       "Test custom input file config",
			inputFile:  "new_input.txt",
			outputFile: "",
			expected: Config{
				InputFile:  "new_input.txt",
				OutputFile: "output.txt",
				LogLevel:   "ERROR",
			},
		},
		{
			name:       "Test custom output file config",
			inputFile:  "",
			outputFile: "new_output.txt",
			expected: Config{
				InputFile:  "input.txt",
				OutputFile: "new_output.txt",
				LogLevel:   "ERROR",
			},
		},
		{
			name:       "Test custom log level",
			inputFile:  "",
			outputFile: "",
			logLevel:   "DEBUG",
			expected: Config{
				InputFile:  "input.txt",
				OutputFile: "output.txt",
				LogLevel:   "DEBUG",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.inputFile != "" {
				t.Setenv(envInputFile, tc.inputFile)
			}

			if tc.outputFile != "" {
				t.Setenv(envOutputFile, tc.outputFile)
			}

			if tc.logLevel != "" {
				t.Setenv(envLogLevel, tc.logLevel)
			}

			cfg := NewConfig()

			assert.Equal(t, tc.expected, cfg)
		})
	}
}
