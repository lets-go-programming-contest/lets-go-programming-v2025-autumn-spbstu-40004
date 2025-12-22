package config

import (
	"github.com/go-yaml/yaml"
)

type Config struct {
	Environment string `yaml:"environment"`
	Log_level   string `yaml:"log_level"`
}

func GetConfig() (*Config, error) {
	var conf Config
	if err := yaml.Unmarshal(ConfigYaml, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
