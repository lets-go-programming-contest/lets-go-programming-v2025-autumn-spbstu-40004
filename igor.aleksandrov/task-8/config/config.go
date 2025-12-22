package config

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

var (
	ErrSourceNotInitialized = errors.New("config source is not initialized")
	ErrConfigDataEmpty      = errors.New("config data is empty")
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func Load() (*Config, error) {
	raw := getConfigData()
	if len(raw) == 0 {
		return nil, ErrConfigDataEmpty
	}

	var cfg Config
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
