package writer

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/arinaklimova/task-3/internal/models"
)

const dirPerm = 0o755

func WriteJSON(currencies []models.Currency, outputPath string) error {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	output := make([]models.CurrencyOutput, len(currencies))
	for i, currency := range currencies {
		output[i] = models.CurrencyOutput(currency)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(output); err != nil {
		return err
	}

	return nil
}
