package internal

import (
	"encoding/xml"
	"os"
)

type Currencies struct {
	Currencies []CurrencyXML `xml:"Valute"`
}

type CurrencyXML struct {
	numericalCode int     `xml:"NumCode"`
	characterCode string  `xml:"CharCode"`
	value         float32 `xml:"Value"`
}

func ParseXML(filePath string) (*Currencies, error) {
	var currencies Currencies

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(file)

	if err := decoder.Decode(&currencies); err != nil {
		return nil, err
	}

	return &currencies, nil
}
