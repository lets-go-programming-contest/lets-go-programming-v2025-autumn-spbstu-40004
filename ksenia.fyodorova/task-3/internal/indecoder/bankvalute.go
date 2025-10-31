package indecoder

import (
	"strconv"
	"strings"
)

type CurrencyItem struct {
	NumericCode     int     `xml:"NumCode" json:"num_code"`
	SymbolCode      string  `xml:"CharCode" json:"char_code"`
	OriginalValue   string  `xml:"Value" json:"-"`
	ConvertedAmount float64 `xml:"-" json:"value"`
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
