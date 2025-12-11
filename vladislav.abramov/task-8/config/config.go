package config

import "errors"

var ErrUnmarshal = errors.New("can't unmarshal yaml")

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}
