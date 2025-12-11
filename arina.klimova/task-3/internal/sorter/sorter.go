package sorter

import (
	"sort"

	"github.com/arinaklimova/task-3/internal/models"
)

func SortCurrenciesByValueDesc(currencies []models.Currency) {
	if len(currencies) == 0 {
		return
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}
