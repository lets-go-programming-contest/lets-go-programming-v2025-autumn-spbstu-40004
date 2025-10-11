package config_handler

import (
	"flag"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type ConfigurationFile struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func loadConfig() (*ConfigurationFile, error) {
	var configFilePath string

	flag.StringVar(&configFilePath, "config", "none", "Configuration file path")
	flag.Parse()

	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var configFile ConfigurationFile
	if err = yaml.Unmarshal(file, &configFile); err != nil {
		return nil, err
	}

	return &configFile, nil
}
