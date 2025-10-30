package currencyProcessor

import (
	"sort"
	"strings"

	"github.com/ZakirovMS/task-3/internal/codingProcessor"
)

func sortValue(val *codingProcessor.ValCurs) {
	for loc := range val.Valutes {
		val.Valutes[loc].Value = strings.ReplaceAll(strings.TrimSpace(val.Valutes[loc].Value), ",", ".")
	}

	sort.Sort(val)
}
