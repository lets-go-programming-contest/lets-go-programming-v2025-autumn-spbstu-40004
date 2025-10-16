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
	if err == nil {
		file, err = os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
	} else if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
	} else {
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
