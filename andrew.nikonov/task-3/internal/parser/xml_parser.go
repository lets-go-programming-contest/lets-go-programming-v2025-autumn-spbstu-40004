package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func Parse(path string, result any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read XML file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(result)
	if err != nil {
		return fmt.Errorf("failed to parse XML: %w", err)
	}

	return nil
}
