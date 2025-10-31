package parser

import (
	"encoding/xml"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ysffmn/task-3/internal/currency"
	"golang.org/x/text/encoding/charmap"
)

func ParseXMLFile(filePath string) currency.ValCurs {
	file, err := os.Open(filePath)
	if err != nil {
		panic("Failed to open XML file: " + err.Error())
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			panic("Failed to close XML file: " + closeErr.Error())
		}
	}()

	var valCurs currency.ValCurs

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return input, nil
	}

	if err := decoder.Decode(&valCurs); err != nil {
		panic("Failed to decode XML: " + err.Error())
	}

	return valCurs
}

func ConvertToCurrencies(valCurs currency.ValCurs) []currency.Currency {
	currencies := make([]currency.Currency, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		valueStr := strings.ReplaceAll(valute.Value, ",", ".")
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
