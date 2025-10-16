package xml_handling

type CurrencyJSON struct {
	NumericalCode int     `json:"num_code"`
	CharacterCode string  `json:"char_code"`
	Value         float32 `json:"value"`
}

func ConvertXMLStructsToJson(currenciesXML CurrenciesXML) []CurrencyJSON {
	arrayLength := len(currenciesXML.Currencies)

	var currenciesJSON = make([]CurrencyJSON, arrayLength)

	for index := 0; index < arrayLength; index++ {
		currenciesJSON[index] = CurrencyJSON(currenciesXML.Currencies[index])
	}

	return currenciesJSON
}
