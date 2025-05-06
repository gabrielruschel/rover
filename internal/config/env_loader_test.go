package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvLoader_GetString(t *testing.T) {
	testCases := []struct {
		name         string
		envVar       string
		setEnvValue  string
		expected     string
		expectedElem string
	}{
		{
			name:         "Test valid string",
			envVar:       "TEST_STRING",
			setEnvValue:  "Hello, World!",
			expected:     "Hello, World!",
			expectedElem: "Hello, World!",
		},
		{
			name:         "Test empty string",
			envVar:       "TEST_STRING",
			setEnvValue:  "",
			expected:     "",
			expectedElem: "",
		},
		{
			name:         "Test unset environment variable",
			envVar:       "TEST_STRING",
			setEnvValue:  "",
			expected:     "",
			expectedElem: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set the environment variable
			t.Setenv(tc.envVar, tc.setEnvValue)

			var element string
			loader := EnvLoader{}
			result := loader.GetString(&element, tc.envVar)

			assert.Equal(t, tc.expected, result)
			assert.Equal(t, tc.expectedElem, element)
		})
	}
}
