package config

const (
	envInputFile  = "INPUT_FILE"
	envOutputFile = "OUTPUT_FILE"
	envLogLevel   = "LOG_LEVEL"
)

type Config struct {
	InputFile  string
	OutputFile string
	LogLevel   string
}

func NewConfig() Config {
	defaultCfg := Config{
		InputFile:  "input.txt",
		OutputFile: "output.txt",
		LogLevel:   "ERROR",
	}

	loader := EnvLoader{}
	loader.GetString(&defaultCfg.InputFile, envInputFile)
	loader.GetString(&defaultCfg.OutputFile, envOutputFile)
	loader.GetString(&defaultCfg.LogLevel, envLogLevel)

	return defaultCfg
}
