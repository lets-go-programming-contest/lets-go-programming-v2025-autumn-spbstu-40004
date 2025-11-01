package indecoder

import (
	"fmt"
	"strconv"
	"strings"
)

type bankValute struct {
	NumCode    int     `json:"num_code"  xml:"NumCode"`
	CharCode   string  `json:"char_code" xml:"CharCode"`
	Value      string  `json:"-"         xml:"Value"`
	FloatValue float64 `json:"value"     xml:"-"`
}

func (val *bankValute) convertFloatValue() error {
	val.Value = strings.ReplaceAll(val.Value, ",", ".")

	floatValue, err := strconv.ParseFloat(val.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to convert float value: %w", err)
	}

	val.FloatValue = floatValue

	return nil
}
