package confdecoder

import (
	"os"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	SourcePath string `yaml:"input-file"`
	TargetPath string `yaml:"output-file"`
}

func ParseConfiguration(configPath string) (*AppConfig, error) {
	fileData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config AppConfig
	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		return nil, err
	}

	if config.SourcePath == "" || config.TargetPath == "" {
		return nil, os.ErrInvalid
	}

	return &config, nil
}
