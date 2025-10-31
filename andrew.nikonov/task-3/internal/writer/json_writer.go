package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const directoryPermission = 0o755

func Write(path string, data any) error {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, directoryPermission)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}
