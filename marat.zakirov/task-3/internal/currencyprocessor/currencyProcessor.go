package currencyprocessor

import (
	"encoding/xml"
	"sort"
	"strconv"
	"strings"
)

type PathHolder struct {
	InPath  string `yaml:"input-file"`
	OutPath string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	ValueStr string  `json:"-"         xml:"Value"`
	ValueFlt float64 `json:"value"     xml:"-"`
}

func (val ValCurs) Len() int {
	return len(val.Valutes)
}

func (val ValCurs) Swap(lhs, rhs int) {
	val.Valutes[lhs], val.Valutes[rhs] = val.Valutes[rhs], val.Valutes[lhs]
}

func (val ValCurs) Less(lhs, rhs int) bool {
	return val.Valutes[lhs].ValueFlt < val.Valutes[rhs].ValueFlt
}

func SortValue(val *ValCurs) {
	var err error

	for loc := range val.Valutes {
		correctStr := strings.ReplaceAll(strings.TrimSpace(val.Valutes[loc].ValueStr), ",", ".")

		val.Valutes[loc].ValueFlt, err = strconv.ParseFloat(correctStr, 64)
		if err != nil {
			panic(err)
		}
	}

	sort.Sort(sort.Reverse(val))
}
