package converter

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"

	"github.com/widgeiw/task-3/internal/models"
)

func ParseXML(filePath string) (*models.ValCurs, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read xml file %w", err)
	}

	reader, err := charset.NewReader(bytes.NewReader(data), "")
	if err != nil {
		return nil, fmt.Errorf("failed to create charset reader %w", err)
	}

	var valCurs models.ValCurs
	decoder := xml.NewDecoder(reader)

	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall xml %w", err)
	}

	return &valCurs, nil
}
