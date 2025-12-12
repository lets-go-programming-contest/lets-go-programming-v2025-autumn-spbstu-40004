package converter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/widgeiw/task-3/internal/models"
)

const (
	dir = 0o755
)

func WriteJSON(currencies []models.Currency, filePath string) error {
	err := os.MkdirAll(filepath.Dir(filePath), dir)
	if err != nil {
		return fmt.Errorf("failed to create directory %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %w", err)
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			panic(err)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(currencies)
	if err != nil {
		return fmt.Errorf("failed to encode json %w", err)
	}

	return nil
}
