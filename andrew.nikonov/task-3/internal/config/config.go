package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		panic("Failed to open config: " + err.Error())
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic("Failed to read config: " + err.Error())
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic("Failed to parse config: " + err.Error())
	}

	return config
}
