package parser

import (
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"

	"github.com/arinaklimova/task-3/internal/models"
)

func ParseXML(filePath string) (*models.ValCurs, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var valCurs models.ValCurs
	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&valCurs); err != nil {
		return nil, err
	}

	return &valCurs, nil
}
