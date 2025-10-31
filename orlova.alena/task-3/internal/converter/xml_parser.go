package converter

import (
	"encoding/xml"
	"os"

	"task-3/internal/models"
)

func ParseXML(filePath string) *models.ValCurs {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic("failed to read xml file" + err.Error())
	}

	var valCurs models.ValCurs

	err = xml.Unmarshal(data, &valCurs)
	if err != nil {
		panic("failed to unmarshall xml" + err.Error())
	}

	return &valCurs
}
