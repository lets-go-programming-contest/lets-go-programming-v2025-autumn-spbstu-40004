package main

import (
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(file string) (Config, error) {
	var config Config

	data, err := os.ReadFile(file)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)

	return config, err
}
