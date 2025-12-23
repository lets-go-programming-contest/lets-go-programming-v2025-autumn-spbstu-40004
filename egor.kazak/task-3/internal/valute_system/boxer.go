package valsys

import (
	"encoding/json"
	"errors"
	"sort"
)

var errMarsJSON = errors.New("cant marshall json")

func CreateJSON(curs *ValCurs) ([]byte, error) {
	cursTemp := make([]Valute, 0, len(curs.Valutes))

	for _, value := range curs.Valutes {
		valTemp := Valute{
			NumCode:  value.NumCode,
			CharCode: value.CharCode,
			Value:    value.Value,
		}

		cursTemp = append(cursTemp, valTemp)
	}

	sort.Slice(cursTemp, func(i, j int) bool {
		return cursTemp[i].Value > cursTemp[j].Value
	})

	jsonData, err := json.MarshalIndent(cursTemp, "", "  ")
	if err != nil {
		return nil, errMarsJSON
	}

	return jsonData, nil
}
