package ioutils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mkryloff/task-3/internal/currencies"
)

const (
	aRW = 0o755
)

func WriteJSONOutput(currencies []currencies.Valute, outputPath string) error {
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, aRW); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			_ = closeErr
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}
