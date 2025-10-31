package processor

import (
	"sort"

	"github.com/ysffmn/task-3/internal/currency"
)

func SortCurrenciesDesc(currencies []currency.Currency) []currency.Currency {
	sorted := make([]currency.Currency, len(currencies))
	copy(sorted, currencies)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	return sorted
}
