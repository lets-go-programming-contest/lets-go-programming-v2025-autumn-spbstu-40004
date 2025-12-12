package sort

import (
	"sort"

	"spbstu.ru/nadia.voronina/task-3/pkg/valute"
)

func SortDescendingByValue(valCurs []valute.Valute) {
	sort.Slice(valCurs, func(i, j int) bool {
		valI, errI := valute.ParseValue(valCurs[i].Value)
		valJ, errJ := valute.ParseValue(valCurs[j].Value)

		if errI != nil || errJ != nil {
			return false
		}

		return valI > valJ
	})
}
