package configparser

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

type configRecord struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ParseConfig(configPath string) (*configRecord, error) {
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var conRec configRecord

	err = yaml.Unmarshal(configFile, &conRec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml file: %w", err)
	}

	return &conRec, nil
}
