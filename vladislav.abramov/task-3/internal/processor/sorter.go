package processor

import (
	"sort"

	"github.com/15446-rus75/task-3/internal/types"
)

func SortCurrenciesByValue(valutes []types.Valute) []types.CurrencyOutput {
	output := make([]types.CurrencyOutput, len(valutes))

	for i, valute := range valutes {
		output[i] = types.CurrencyOutput{
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
