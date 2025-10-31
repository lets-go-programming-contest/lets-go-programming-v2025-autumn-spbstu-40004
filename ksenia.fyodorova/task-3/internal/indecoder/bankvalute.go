package indecoder

import (
	"strconv"
	"strings"
)

type CurrencyItem struct {
	NumericCode     int     `xml:"NumCode"`
	SymbolCode      string  `xml:"CharCode"`
	OriginalValue   string  `xml:"Value"`
	ConvertedAmount float64 `xml:"-"`
}

func (ci *CurrencyItem) TransformValue() error {
	formattedValue := strings.Replace(ci.OriginalValue, ",", ".", 1)
	parsedValue, err := strconv.ParseFloat(formattedValue, 64)
	if err != nil {
		return err
	}

	ci.ConvertedAmount = parsedValue
	return nil
}
