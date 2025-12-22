package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func Parse() (*Config, error) {
	var config Config

	err := yaml.Unmarshal(configYamlData, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	return &config, nil
}
