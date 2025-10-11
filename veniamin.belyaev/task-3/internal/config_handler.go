package internal

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type configurationFile struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(configFilePath string) (*configurationFile, error) {
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var configFile configurationFile
	if err = yaml.Unmarshal(file, &configFile); err != nil {
		return nil, err
	}

	return &configFile, nil
}
