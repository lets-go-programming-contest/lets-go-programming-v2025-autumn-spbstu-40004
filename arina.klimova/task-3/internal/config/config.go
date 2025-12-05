package config

import (
	"flag"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig() *Config {
	configPath := LoadConfigPath()

	file, err := os.Open(configPath)
	if err != nil {
		panic("failed to open config file: " + err.Error())
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic("failed to close config file: " + err.Error())
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		panic("failed to read config file: " + err.Error())
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic("failed to parse config file: " + err.Error())
	}

	config.validate()

	return &config
}

func (c *Config) validate() {
	if c.InputFile == "" {
		panic("input-file is required in config")
	}

	if c.OutputFile == "" {
		panic("output-file is required in config")
	}
}

func LoadConfigPath() string {
	var configPath string

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if configPath == "" {
		panic("config flag is required")
	}

	return configPath
}
