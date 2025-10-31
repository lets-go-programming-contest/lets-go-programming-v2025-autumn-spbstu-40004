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
	parsedValue, parseErr := strconv.ParseFloat(formattedValue, 64)
	if parseErr != nil {
		return parseErr
	}

	ci.ConvertedAmount = parsedValue
	return nil
}
