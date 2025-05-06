package config

import "os"

type EnvLoader struct{}

func (el EnvLoader) GetString(element *string, envName string) string {
	result := os.Getenv(envName)
	if result != "" {
		*element = result
	}
	return *element
}
