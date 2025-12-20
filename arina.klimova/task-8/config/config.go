package config

import (
	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	data := getConfigData()

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
