package config

import (
	_ "embed"
	"errors"
	"gopkg.in/yaml.v3"
)

var devConfig []byte

var ErrUnmarshal = errors.New("can't unmarshal yaml")

func Load() (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(devConfig, &cfg); err != nil {
		return nil, ErrUnmarshal
	}
	return &cfg, nil
}
