package writer

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/ysffmn/task-3/internal/currency"
)

func WriteJSONToFile(currencies []currency.Currency, outputPath string) {
	outputDir := filepath.Dir(outputPath)
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		panic("Failed to create output directory: " + err.Error())
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic("Failed to create output file: " + err.Error())
	}
	defer outputFile.Close()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(currencies)
	if err != nil {
		panic("Failed to encode JSON: " + err.Error())
	}
}
