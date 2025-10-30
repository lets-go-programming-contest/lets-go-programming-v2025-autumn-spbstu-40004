package currencyProcessor

import (
	"strconv"
	"strings"
)

func parseValue(valueStr string) float64 {
	cleaned := strings.ReplaceAll(strings.TrimSpace(valueStr), ",", ".")
	val, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return 0
	}

	return val
}
