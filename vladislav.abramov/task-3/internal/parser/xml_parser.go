package parser

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"

	"internal/types"
	"golang.org/x/net/html/charset"
)

func ParseCurrencyData(filePath string) (*types.CurrencyData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	contentWithDots := bytes.ReplaceAll(content, []byte(","), []byte("."))

	decoder := xml.NewDecoder(bytes.NewReader(contentWithDots))
	decoder.CharsetReader = charset.NewReaderLabel

	var data types.CurrencyData
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
