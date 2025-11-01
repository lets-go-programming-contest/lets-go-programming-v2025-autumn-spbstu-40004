package centralbank

import (
	"encoding/json"
	"fmt"
	"os"
)

const outputFilePerm = 0o644

func SaveCurrenciesToJSON(currencies []Currency, outputFile string) error {
	jsonData, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal to JSON: %w", err)
	}

	err = os.WriteFile(outputFile, jsonData, outputFilePerm)
	if err != nil {
		return fmt.Errorf("cannot write to file %s: %w", outputFile, err)
	}

	return nil
}
