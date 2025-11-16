package processing

import (
	"sort"

	"github.com/mkryloff/task-3/internal/currencies"
)

func SortCurrenciesByValue(valutes []currencies.Valute) []currencies.CurrencyOutput {
	output := make([]currencies.CurrencyOutput, len(valutes))

	for i, valute := range valutes {
		output[i] = currencies.CurrencyOutput{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    valute.Value,
		}
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Value > output[j].Value
	})

	return output
}
