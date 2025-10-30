package ioutils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfiguration(configPath string) (*Configuration, error) {
	fileData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Configuration
	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
