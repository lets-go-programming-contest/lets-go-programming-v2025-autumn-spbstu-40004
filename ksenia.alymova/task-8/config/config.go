package config

import (
	"fmt"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func GetConfig() (*Config, error) {
	var conf Config
	if err := yaml.Unmarshal(ConfigYaml, &conf); err != nil {
		return nil, fmt.Errorf("config file error: %w", err)
	}

	return &conf, nil
}
