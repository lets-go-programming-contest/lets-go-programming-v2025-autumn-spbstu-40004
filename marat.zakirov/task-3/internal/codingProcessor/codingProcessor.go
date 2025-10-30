package codingProcessor

import (
	"strconv"

	"github.com/ZakirovMS/task-3/internal/currencyProcessor"
)

type PathHolder struct {
	InPath  string `yaml:"input-file"`
	OutPath string `yaml:"output-file"`
}

type jsonCurs struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func convertValute(valute currencyProcessor.Valute) jsonCurs {
	var result jsonCurs
	var err error

	result.NumCode, err = strconv.Atoi(valute.NumCode)
	if err != nil {
		panic("Some errors in NumCode conversion")
	}

	result.CharCode = valute.CharCode

	result.Value, err = strconv.ParseFloat(valute.Value, 64)
	if err != nil {
		panic("Some errors in Value conversion")
	}

	return result
}

func ConvertXmlToJson(val currencyProcessor.ValCurs) []jsonCurs {
	result := make([]jsonCurs, 0, len(val.Valutes))

	for _, valute := range val.Valutes {
		converted := convertValute(valute)
		result = append(result, converted)
	}

	return result
}
