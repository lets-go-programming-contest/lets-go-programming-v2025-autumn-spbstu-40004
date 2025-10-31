package converter

import (
	"encoding/json"
	"os"
	"path/filepath"

	"task-3/internal/models"
)

func WriteJSON(currencies []models.Currency, filePath string) error {
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(currencies)
}
