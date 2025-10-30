package ioutils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/15446-rus75/task3/internal/types"
)

func WriteJSONOutput(currencies []types.CurrencyOutput, outputPath string) error {
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(currencies)
}
