package sorter

import (
	"sort"
	"strconv"
	"strings"

	"github.com/shycoshy/task-3/internal/domain"
)

func Sort(valutes []domain.Valute) []domain.Currency {
	results := make([]domain.Currency, len(valutes))

	for i, v := range valutes {
		value, _ := strconv.ParseFloat(strings.Replace(v.Value, ",", ".", 1), 64)
		results[i] = domain.Currency{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    value,
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})

	return results
}
