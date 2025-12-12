package converter

import (
	"sort"
	"strconv"
	"strings"

	"github.com/widgeiw/task-3/internal/models"
)

func ConvertToCurrencies(valCurs *models.ValCurs) []models.Currency {
	currencies := make([]models.Currency, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		value, err := strconv.ParseFloat(strings.Replace(valute.Value, ",", ".", 1), 64)
		if err != nil {
			continue
		}

		currencies = append(currencies, models.Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies
}
