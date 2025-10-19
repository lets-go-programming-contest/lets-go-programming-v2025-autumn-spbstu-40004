package currency

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Write(path string, currencies *Currencies) error {
	const allReadWriteMode = os.FileMode(0o666)

	var file *os.File

	_, err := os.Stat(path)

	switch {
	case err == nil:
		file, err = os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
	case errors.Is(err, os.ErrNotExist):
		err := os.MkdirAll(filepath.Dir(path), allReadWriteMode)
		if err != nil {
			return fmt.Errorf("failed to make directory: %w", err)
		}

		file, err = os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
	default:
		return fmt.Errorf("failed to get file info: %w", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(currencies.Data)
	if err != nil {
		return fmt.Errorf("failed to encode to json file: %w", err)
	}

	return nil
}
