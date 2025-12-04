package outcoder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lolnyok/task-3/internal/indecoder"
)

const (
	dirPerm  = 0o755
	filePerm = 0o600
)

func SaveCurrencyData(outputPath string, data indecoder.CurrencyCollection) error {
	dir := filepath.Dir(outputPath)

	err := os.MkdirAll(dir, dirPerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	jsonData, err := json.MarshalIndent(data.Items, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = os.WriteFile(outputPath, jsonData, filePerm)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
