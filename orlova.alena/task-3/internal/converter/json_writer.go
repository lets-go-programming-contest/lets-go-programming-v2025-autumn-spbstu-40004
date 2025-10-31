package converter

import (
	"encoding/json"
	"os"
	"path/filepath"

	"task-3/internal/models"
)

const (
	dir = 0o755
)

func WriteJSON(currencies []models.Currency, filePath string) error {
	err := os.MkdirAll(filepath.Dir(filePath), dir)
	if err != nil {
		panic("failed to create directory" + err.Error())
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic("failed to create file" + err.Error())
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			_ = closeErr
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(currencies)
	if err != nil {
		panic("failed to encode json" + err.Error())
	}

	return nil
}
