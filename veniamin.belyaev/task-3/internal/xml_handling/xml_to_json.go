package xmlhandling

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type CurrencyJSON struct {
	NumericalCode int     `json:"num_code"`
	CharacterCode string  `json:"char_code"`
	Value         float32 `json:"value"`
}

func ConvertXMLStructsToJSON(currenciesXML CurrenciesXML) ([]CurrencyJSON, error) {
	arrayLength := len(currenciesXML.Currencies)

	currenciesJSON := make([]CurrencyJSON, arrayLength)

	for index := range arrayLength {
		valueString := currenciesXML.Currencies[index].Value

		valueString = strings.Replace(valueString, ",", ".", 1)

		valueFloat, err := strconv.ParseFloat(valueString, 32)
		if err != nil {
			return nil, fmt.Errorf("strconv: %w", err)
		}

		currenciesJSON[index] = CurrencyJSON{
			NumericalCode: currenciesXML.Currencies[index].NumericalCode,
			CharacterCode: currenciesXML.Currencies[index].CharacterCode,
			Value:         float32(valueFloat),
		}
	}

	sort.Slice(currenciesJSON, func(i, j int) bool {
		return currenciesJSON[i].Value > currenciesJSON[j].Value
	})

	return currenciesJSON, nil
}
