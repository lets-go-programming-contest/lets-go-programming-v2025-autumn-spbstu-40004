package main

import (
	"sort"
)

func SortDescendingByValue(valCurs []ValuteJSON) {
	sort.Slice(valCurs, func(i, j int) bool {
		return valCurs[i].Value > valCurs[j].Value
	})
}
