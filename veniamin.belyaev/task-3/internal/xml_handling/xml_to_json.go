package xml_handling

import (
	"strconv"
	"strings"
)

type CurrencyJSON struct {
	NumericalCode int     `json:"num_code"`
	CharacterCode string  `json:"char_code"`
	Value         float32 `json:"value"`
}

func ConvertXMLStructsToJson(currenciesXML CurrenciesXML) ([]CurrencyJSON, error) {
	arrayLength := len(currenciesXML.Currencies)

	var currenciesJSON = make([]CurrencyJSON, arrayLength)

	for index := 0; index < arrayLength; index++ {
		valueString := currenciesXML.Currencies[index].Value

		valueString = strings.Replace(valueString, ",", ".", 1)

		valueFloat, err := strconv.ParseFloat(valueString, 32)
		if err != nil {
			return nil, err
		}

		currenciesJSON[index] = CurrencyJSON{
			NumericalCode: currenciesXML.Currencies[index].NumericalCode,
			CharacterCode: currenciesXML.Currencies[index].CharacterCode,
			Value:         float32(valueFloat),
		}
	}

	return currenciesJSON, nil
}
