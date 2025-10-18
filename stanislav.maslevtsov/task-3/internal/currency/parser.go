package currency

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
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
		return nil, fmt.Errorf("failed to open xml currencies file: %w", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	var (
		xmlCurs xmlCurrencies
		decoder = xml.NewDecoder(file)
	)

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	err = decoder.Decode(&xmlCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml currencies file: %w", err)
	}

	curs, err := convertXMLToCurrency(&xmlCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to convert xml currency: %w", err)
	}

	return curs, nil
}

func convertXMLToCurrency(xmlCurs *xmlCurrencies) (*Currencies, error) {
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
