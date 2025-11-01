package confighandling

import (
	"errors"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

var errYamlKeyNotFound = errors.New("did not find expected key")

type AppConfig struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfiguration(configPath string) (*AppConfig, error) {
	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("config handler i/o: %w", err)
	}

	var configData AppConfig
	if err = yaml.Unmarshal(fileContent, &configData); err != nil {
		return nil, fmt.Errorf("yaml: %w", err)
	}

	if configData.InputFile == "" || configData.OutputFile == "" {
		return nil, errYamlKeyNotFound
	}

	return &configData, nil
}
