package config_handler

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type ConfigurationFile struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(configFilePath string) (*ConfigurationFile, error) {
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var configFile ConfigurationFile
	if err = yaml.Unmarshal(file, &configFile); err != nil {
		return nil, err
	}

	return &configFile, nil
}
