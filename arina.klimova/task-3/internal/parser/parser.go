package parser

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/arinaklimova/task-3/internal/models"
	"golang.org/x/net/html/charset"
)

func ParseXML(filePath string) (*models.ValCurs, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open XML file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic("failed to close XML file: " + err.Error())
		}
	}()

	var valCurs models.ValCurs

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	return &valCurs, nil
}
