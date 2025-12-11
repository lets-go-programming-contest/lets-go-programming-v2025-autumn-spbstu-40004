package config

import (
	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func Load() *Config {
	return loadConfig("prod.yaml")
}

func loadConfig(filename string) *Config {
	data, err := prodYAML.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}
