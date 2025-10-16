package xmlhandling

import (
	"encoding/xml"
	"fmt"
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
		return nil, fmt.Errorf("i/o: %s", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err.Error())
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&currencies); err != nil {
		return nil, fmt.Errorf("decoder: %s", err)
	}

	return &currencies, nil
}
