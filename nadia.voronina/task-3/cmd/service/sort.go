package main

import (
	"sort"
)

func SortDescendingByValue(valCurs []ValuteJson) {
	sort.Slice(valCurs, func(i, j int) bool {
		return valCurs[i].Value > valCurs[j].Value
	})
}
