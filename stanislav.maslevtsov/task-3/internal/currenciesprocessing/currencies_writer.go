package currenciesprocessing

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func WriteCurrencies(path string, currencies *Currencies) error {
	var file *os.File

	_, err := os.Stat(path)

	switch {
	case err == nil:
		file, err = os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
	case errors.Is(err, os.ErrNotExist):
		file, err = os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
	default:
		return fmt.Errorf("failed to get file info: %w", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(currencies.Data)
	if err != nil {
		return fmt.Errorf("failed to encode to file: %w", err)
	}

	return nil
}
