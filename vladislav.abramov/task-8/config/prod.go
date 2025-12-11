package config

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

var prodConfig []byte

func Load() (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(prodConfig, &cfg); err != nil {
		return nil, ErrUnmarshal
	}
	return &cfg, nil
}
