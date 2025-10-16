package xml_handling

import (
	"encoding/xml"
	"os"
)

type CurrencyXML struct {
	NumericalCode int     `xml:"NumCode"`
	CharacterCode string  `xml:"CharCode"`
	Value         float32 `xml:"Value"`
}

type CurrenciesXML struct {
	Currencies []CurrencyXML `xml:"Valute"`
}

func ParseXML(filePath string) (*CurrenciesXML, error) {
	var currencies CurrenciesXML

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
