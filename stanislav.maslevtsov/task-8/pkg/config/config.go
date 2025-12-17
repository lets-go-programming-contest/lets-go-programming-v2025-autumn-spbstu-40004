package config

import (
	"fmt"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func Parse() (*Config, error) {
	var conf Config

	if err := yaml.Unmarshal(configData, &conf); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &conf, nil
}
