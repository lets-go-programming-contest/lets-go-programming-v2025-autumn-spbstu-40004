package writer

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/shycoshy/task-3/internal/domain"
)

func Save(data []domain.Currency, path string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
