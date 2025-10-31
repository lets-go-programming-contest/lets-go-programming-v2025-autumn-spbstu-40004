package parser

import (
	"encoding/xml"
	"os"
	"strconv"
	"strings"

	"github.com/ysffmn/task-3/internal/currency"
)

func ParseXMLFile(filePath string) currency.ValCurs {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic("Failed to read input file: " + err.Error())
	}

	var valCurs currency.ValCurs
	err = xml.Unmarshal(data, &valCurs)
	if err != nil {
		panic("Failed to parse XML: " + err.Error())
	}

	return valCurs
}

func ConvertToCurrencies(valCurs currency.ValCurs) []currency.Currency {
	currencies := make([]currency.Currency, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		valueStr := strings.Replace(valute.Value, ",", ".", -1)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			panic("Failed to parse currency value: " + err.Error())
		}

		numCode, err := strconv.Atoi(valute.NumCode)
		if err != nil {
			panic("Failed to parse numeric code: " + err.Error())
		}

		currencies = append(currencies, currency.Currency{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	return currencies
}
