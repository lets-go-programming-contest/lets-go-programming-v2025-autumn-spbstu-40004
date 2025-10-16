package configparser

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

type ConfigRecord struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Parse(path string) (*ConfigRecord, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	var (
		conRec  ConfigRecord
		decoder = yaml.NewDecoder(file)
	)

	err = decoder.Decode(&conRec)
	if err != nil {
		return nil, fmt.Errorf("failed to decode yaml file: %w", err)
	}

	return &conRec, nil
}
