package config

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

type sourceFunc func() []byte

var currentSource sourceFunc

func Load() (*Config, error) {
	if currentSource == nil {
		return nil, errors.New("config source is not initialized")
	}

	raw := currentSource()
	if len(raw) == 0 {
		return nil, errors.New("config data is empty")
	}

	var cfg Config
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, fmt.Errorf("parsing error: %w", err)
	}

	return &cfg, nil
}
