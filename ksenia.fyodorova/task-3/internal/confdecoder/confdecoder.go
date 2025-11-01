package confdecoder

import (
	"fmt"
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
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var config AppConfig

	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if config.SourcePath == "" || config.TargetPath == "" {
		return nil, os.ErrInvalid
	}

	return &config, nil
}
