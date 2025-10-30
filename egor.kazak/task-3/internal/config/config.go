package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig() (*Config, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if configPath == "" {
		panic("config path is required")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	return &cfg, nil
}