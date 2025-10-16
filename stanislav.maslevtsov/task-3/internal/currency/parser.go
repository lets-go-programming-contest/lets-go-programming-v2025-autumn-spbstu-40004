package currency

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type xmlCurrency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type xmlCurrencies struct {
	Data []*xmlCurrency `xml:"Valute"`
}

func Parse(path string) (*Currencies, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	var (
		xmlCurs xmlCurrencies
		decoder = xml.NewDecoder(file)
	)

	err = decoder.Decode(&xmlCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml file: %w", err)
	}

	var curs Currencies

	for xmlCurIdx := range xmlCurs.Data {
		strValue := strings.Replace(xmlCurs.Data[xmlCurIdx].Value, ",", ".", 1)

		floatValue, err := strconv.ParseFloat(strValue, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse currency value: %w", err)
		}

		curs.Data = append(curs.Data, &Currency{
			NumCode:  xmlCurs.Data[xmlCurIdx].NumCode,
			CharCode: xmlCurs.Data[xmlCurIdx].CharCode,
			Value:    float32(floatValue),
		})
	}

	return &curs, nil
}
