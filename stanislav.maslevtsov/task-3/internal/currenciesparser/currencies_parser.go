package currenciesparser

import (
	"fmt"
	"os"

	"encoding/xml"
)

type Currency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currencies struct {
	Сurncs []Currency `xml:"Valute"`
}

func ParseCurrencies(path string) (*Currencies, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	var currencies Currencies
	decoder := xml.NewDecoder(file)

	err = decoder.Decode(&currencies)
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml file: %w", err)
	}

	return &currencies, nil
}
