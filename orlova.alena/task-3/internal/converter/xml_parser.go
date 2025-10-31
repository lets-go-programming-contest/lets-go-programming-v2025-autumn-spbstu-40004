package converter

import (
	"encoding/xml"
	"fmt"
	"os"

	"task-3/internal/models"
)

func ParseXML(filePath string) (*models.ValCurs, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read xml file %w", err)
	}

	var valCurs models.ValCurs

	err = xml.Unmarshal(data, &valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall xml %w", err)
	}

	return &valCurs, nil
}
