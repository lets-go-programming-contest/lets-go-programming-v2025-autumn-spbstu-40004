package indecoder

import (
	"fmt"
	"strconv"
	"strings"
)

type CurrencyItem struct {
	NumericCode     int     `json:"num_code"  xml:"NumCode"`
	SymbolCode      string  `json:"char_code" xml:"CharCode"`
	OriginalValue   string  `json:"-"         xml:"Value"`
	ConvertedAmount float64 `json:"value"     xml:"-"`
}

func (ci *CurrencyItem) TransformValue() error {
	formattedValue := strings.Replace(ci.OriginalValue, ",", ".", 1)

	parsedValue, err := strconv.ParseFloat(formattedValue, 64)
	if err != nil {
		return fmt.Errorf("failed to parse currency value: %w", err)
	}

	ci.ConvertedAmount = parsedValue

	return nil
}
