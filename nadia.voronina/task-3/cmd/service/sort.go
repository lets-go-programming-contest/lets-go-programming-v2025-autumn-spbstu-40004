package main

import (
	"sort"
)

func SortDescendingByValue(valCurs *ValCurs) {
	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})
}
