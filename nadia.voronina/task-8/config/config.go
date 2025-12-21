package config

import (
	"github.com/go-yaml/yaml"
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(ConfigYaml, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
