package converter

import (
	"encoding/xml"
	"os"

	"task-3/internal/models"
)

func ParseXML(filePath string) (*models.ValCurs, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var valCurs models.ValCurs
	err = xml.Unmarshal(data, &valCurs)
	if err != nil {
		return nil, err
	}

	return &valCurs, nil
}
