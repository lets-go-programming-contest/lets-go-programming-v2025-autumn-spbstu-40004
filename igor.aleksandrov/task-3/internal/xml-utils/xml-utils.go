package xmlutils

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	Valutes []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  string `xml:"NumCode"  json:"-"`
	CharCode string `xml:"CharCode" json:"char_code"`
	ValueStr string `xml:"Value"    json:"-"`
	NumericalCode int     `xml:"-" json:"num_code"`
	Value         float64 `xml:"-" json:"value"`
}

func ParseXML(filePath string) ([]Currency, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	transformedData := strings.ReplaceAll(string(data), ",", ".")

	var valCurs ValCurs

	decoder := xml.NewDecoder(strings.NewReader(transformedData))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML data: %w", err)
	}

	for i := range valCurs.Valutes {
		currency := &valCurs.Valutes[i]

		value, err := strconv.ParseFloat(currency.ValueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value '%s' to float: %w", currency.ValueStr, err)
		}

		currency.Value = value

		numCode := 0

		if currency.NumCode != "" {
			convertedCode, err := strconv.Atoi(currency.NumCode)
			if err != nil {
				return nil, fmt.Errorf("failed to parse num_code '%s' to int: %w", currency.NumCode, err)
			}

			numCode = convertedCode
		}

		currency.NumericalCode = numCode
	}

	return valCurs.Valutes, nil
}
