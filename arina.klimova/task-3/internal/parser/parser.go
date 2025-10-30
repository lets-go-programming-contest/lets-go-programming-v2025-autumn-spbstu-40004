package parser

import (
	"encoding/xml"
	"fmt"
	"os"

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
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, err
	}

	if len(valCurs.Currencies) == 0 {
		return nil, fmt.Errorf("XML file contains no currency data")
	}

	return &valCurs, nil
}
