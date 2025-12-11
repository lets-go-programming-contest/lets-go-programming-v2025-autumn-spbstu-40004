package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/arinaklimova/task-3/internal/models"
)

const dirPerm = 0o755

func WriteJSON(currencies []models.Currency, outputPath string) error {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic("failed to close JSON file: " + err.Error())
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	return nil
}
