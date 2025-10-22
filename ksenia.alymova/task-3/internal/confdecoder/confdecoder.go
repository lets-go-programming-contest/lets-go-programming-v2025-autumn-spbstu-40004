package confdecoder

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

var errConfig = errors.New("incorrect configuration")

type configFile struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ConfigProcess(flagConfig *string) (configFile, error) {
	var config configFile

	configByte, err := os.ReadFile(*flagConfig)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(configByte, &config)

	if config.InputFile == "" || config.OutputFile == "" {
		return config, errConfig
	}

	return config, err
}
