package writer

import (
	"encoding/json"
	"os"

	"github.com/arinaklimova/task-3/internal/models"
)

func WriteJSON(currencies []models.Currency, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	output := make([]models.CurrencyOutput, len(currencies))
	for i, currency := range currencies {
		output[i] = models.CurrencyOutput{
			NumCode:  currency.NumCode,
			CharCode: currency.CharCode,
			Value:    currency.Value,
		}
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(output); err != nil {
		return err
	}

	return nil
}
