package config

import (
	"fmt"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

type ConfigError struct {
	Message string
	Err     error
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("config error: %s: %v", e.Message, e.Err)
}

func (e *ConfigError) Unwrap() error {
	return e.Err
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(ConfigYaml, &cfg); err != nil {
		return nil, &ConfigError{
			Message: "failed to unmarshal config",
			Err:     err,
		}
	}

	return &cfg, nil
}
