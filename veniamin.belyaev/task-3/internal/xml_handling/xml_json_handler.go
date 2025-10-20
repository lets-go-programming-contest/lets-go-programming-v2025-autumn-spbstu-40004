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
	NumericalCode int    `xml:"NumCode" json:"num_code"`
	CharacterCode string `xml:"CharCode" json:"char_code"`
	Value         string `xml:"Value" json:"value"`
}

type CurrenciesXML struct {
	Currencies []Currency `xml:"Valute"`
}

func ParseXML(filePath string) (*CurrenciesXML, error) {
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

	return &currencies, nil
}

func ConvertXMLStructsToJSONFormat(currenciesXML CurrenciesXML) ([]Currency, error) {
	arrayLength := len(currenciesXML.Currencies)

	currenciesJSON := make([]Currency, arrayLength)

	for index := range arrayLength {
		valueString := currenciesXML.Currencies[index].Value

		valueString = strings.Replace(valueString, ",", ".", 1)

		currenciesJSON[index] = Currency{
			NumericalCode: currenciesXML.Currencies[index].NumericalCode,
			CharacterCode: currenciesXML.Currencies[index].CharacterCode,
			Value:         valueString,
		}
	}

	sort.Slice(currenciesJSON, func(i, j int) bool {
		iFloat, err := strconv.ParseFloat(currenciesJSON[i].Value, 64)
		if err != nil {
			panic(err.Error())
		}

		jFloat, err := strconv.ParseFloat(currenciesJSON[j].Value, 64)
		if err != nil {
			panic(err.Error())
		}

		return iFloat > jFloat
	})

	return currenciesJSON, nil
}
