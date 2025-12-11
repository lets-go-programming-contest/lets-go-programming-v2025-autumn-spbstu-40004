package config

import (
	"embed"
	"gopkg.in/yaml.v3"
)

var configFS embed.FS

type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func Load() *Config {
	filename := "prod.yaml"

	return loadConfig(filename)
}

func loadConfig(filename string) *Config {
	data, err := configFS.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}
