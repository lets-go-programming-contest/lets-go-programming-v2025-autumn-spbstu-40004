package confdecoder

import (
	"errors"
	"fmt"
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
		return config, fmt.Errorf("failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(configByte, &config)
	if err != nil {
		return config, fmt.Errorf("failed to decode config file: %v", err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		return config, errConfig
	}

	return config, nil
}
