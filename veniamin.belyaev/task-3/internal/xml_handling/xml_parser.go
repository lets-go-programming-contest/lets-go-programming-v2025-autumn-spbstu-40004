package xml_handling

import (
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"
)

type CurrencyXML struct {
	NumericalCode int    `xml:"NumCode"`
	CharacterCode string `xml:"CharCode"`
	Value         string `xml:"Value"`
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
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&currencies); err != nil {
		return nil, err
	}

	return &currencies, nil
}
