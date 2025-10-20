package xmlhandling

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	NumericalCode int     `json:"num_code"  xml:"NumCode"`
	CharacterCode string  `json:"char_code" xml:"CharCode"`
	Value         float64 `json:"value"     xml:"Value"`
}

type CurrencyXMLTemp struct {
	NumericalCode int    `xml:"NumCode"`
	CharacterCode string `xml:"CharCode"`
	Value         string `xml:"Value"`
}

type CurrenciesXML struct {
	Currencies []CurrencyXMLTemp `xml:"Valute"`
}

func ParseXML(filePath string) ([]Currency, error) {
	var currencies CurrenciesXML

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("i/o: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err.Error())
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&currencies); err != nil {
		return nil, fmt.Errorf("decoder: %w", err)
	}

	arrayLength := len(currencies.Currencies)
	currenciesFormatted := make([]Currency, arrayLength)

	for index := range arrayLength {
		valueString := currencies.Currencies[index].Value

		valueString = strings.Replace(valueString, ",", ".", 1)

		valueFloat, err := strconv.ParseFloat(valueString, 64)
		if err != nil {
			return nil, fmt.Errorf("strconv: %w", err)
		}

		currenciesFormatted[index] = Currency{
			NumericalCode: currencies.Currencies[index].NumericalCode,
			CharacterCode: currencies.Currencies[index].CharacterCode,
			Value:         valueFloat,
		}
	}

	sort.Slice(currenciesFormatted, func(i, j int) bool {
		return currenciesFormatted[i].Value > currenciesFormatted[j].Value
	})

	return currenciesFormatted, nil
}
