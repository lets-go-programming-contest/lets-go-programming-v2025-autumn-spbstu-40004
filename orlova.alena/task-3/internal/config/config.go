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

func Load() *Config {
	var cfg Config

	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic("failed to read config: " + err.Error())
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic("failed to parse config: " + err.Error())
	}

	return &cfg
}
