package outcoder

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/lolnyok/task-3/internal/indecoder"
)

func SaveCurrencyData(outputPath string, data indecoder.CurrencyCollection) error {
	dir := filepath.Dir(outputPath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(data.Items, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
