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
	XMLName struct{} `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	XMLName  struct{} `xml:"Valute"`
	NumCode  string   `xml:"NumCode"`
	CharCode string   `xml:"CharCode"`
	ValueStr string   `xml:"Value"`
}

type Currency struct {
	NumericalCode int     `json:"num_code"`
	CharacterCode string  `json:"char_code"`
	Value         float64 `json:"value"`
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

	result := make([]Currency, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		value, err := strconv.ParseFloat(valute.ValueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value '%s' to float: %w", valute.ValueStr, err)
		}

		var numCode int

		if valute.NumCode != "" {
			convertedCode, err := strconv.Atoi(valute.NumCode)
			if err != nil {
				return nil, fmt.Errorf("failed to parse num_code '%s' to int: %w", valute.NumCode, err)
			}

			numCode = convertedCode
		}

		result = append(result, Currency{
			NumericalCode: numCode,
			CharacterCode: valute.CharCode,
			Value:         value,
		})
	}

	return result, nil
}
