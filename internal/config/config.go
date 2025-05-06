package config

const (
	envInputFile  = "INPUT_FILE"
	envOutputFile = "OUTPUT_FILE"
)

type Config struct {
	InputFile  string
	OutputFile string
}

func NewConfig() Config {
	defaultCfg := Config{
		InputFile:  "input.txt",
		OutputFile: "output.txt",
	}

	loader := EnvLoader{}
	loader.GetString(&defaultCfg.InputFile, envInputFile)
	loader.GetString(&defaultCfg.OutputFile, envInputFile)

	return defaultCfg
}
