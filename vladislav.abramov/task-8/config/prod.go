package config

import (
	_ "embed"
	"errors"
	"gopkg.in/yaml.v3"
)

var prodConfig []byte

var ErrUnmarshal = errors.New("can't unmarshal yaml")

func Load() (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(prodConfig, &cfg); err != nil {
		return nil, ErrUnmarshal
	}
	return &cfg, nil
}
