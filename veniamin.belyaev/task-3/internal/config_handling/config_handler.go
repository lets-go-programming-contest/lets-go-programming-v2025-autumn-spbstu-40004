package confighandling

import (
	"errors"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

var errYamlKeyNotFound = errors.New("did not find expected key")

type configurationFile struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(configFilePath string) (*configurationFile, error) {
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("config handler i/o: %w", err)
	}

	var configFile configurationFile
	if err = yaml.Unmarshal(file, &configFile); err != nil {
		return nil, fmt.Errorf("yaml: %w", err)
	}

	if configFile.InputFile == "" || configFile.OutputFile == "" {
		return nil, errYamlKeyNotFound
	}

	return &configFile, nil
}
