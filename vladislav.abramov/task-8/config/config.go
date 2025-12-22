package config

import (
	_ "embed"
	"errors"

	"gopkg.in/yaml.v3"
)

var ErrUnmarshal = errors.New("can't unmarshal yaml")

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

//go:embed prod.yaml
var prodConfig []byte

func loadConfig(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, ErrUnmarshal
	}

	return &cfg, nil
}
