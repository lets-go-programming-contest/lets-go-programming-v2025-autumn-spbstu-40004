package outcoder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lolnyok/task-3/internal/indecoder"
)

func SaveCurrencyData(outputPath string, data indecoder.CurrencyCollection) error {
	dir := filepath.Dir(outputPath)
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	jsonData, err := json.MarshalIndent(data.Items, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = os.WriteFile(outputPath, jsonData, 0o600)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
