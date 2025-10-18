package config

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
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	var (
		conRec  ConfigRecord
		decoder = yaml.NewDecoder(file)
	)

	err = decoder.Decode(&conRec)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return &conRec, nil
}
