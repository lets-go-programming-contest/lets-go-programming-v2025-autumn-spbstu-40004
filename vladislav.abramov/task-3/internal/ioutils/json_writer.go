package ioutils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/15446-rus75/task-3/internal/types"
)

const (
	aRW = 0o755
)

func WriteJSONOutput(currencies []types.CurrencyOutput, outputPath string) error {
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
			panic("failed to close file: " + closeErr.Error())
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}
