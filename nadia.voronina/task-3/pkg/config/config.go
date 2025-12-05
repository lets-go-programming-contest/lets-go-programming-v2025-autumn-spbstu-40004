package config

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type FileReadingError struct {
	File string
	Err  error
}

func (e *FileReadingError) Error() string {
	return fmt.Sprintf("failed to read file %q: %v", e.File, e.Err)
}

type FileUnmarshalError struct {
	File string
	Err  error
}

func (e *FileUnmarshalError) Error() string {
	return fmt.Sprintf("failed to unmarshal yaml from %q: %v", e.File, e.Err)
}

func LoadConfig(file string) (Config, error) {
	var config Config

	data, err := os.ReadFile(file)
	if err != nil {
		return config, &FileReadingError{File: file, Err: err}
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, &FileUnmarshalError{File: file, Err: err}
	}

	return config, nil
}
