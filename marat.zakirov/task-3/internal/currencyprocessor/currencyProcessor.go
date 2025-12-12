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
	NumCode  string  `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	ValueStr string  `xml:"Value" json:"-"`
	ValueFlt float64 `xml:"-" json:"value"`
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
		val.Valutes[loc].ValueFlt, err = strconv.ParseFloat(strings.ReplaceAll(strings.TrimSpace(val.Valutes[loc].ValueStr), ",", "."), 64)
	}

	if err != nil {
		panic(err)
	}

	sort.Sort(sort.Reverse(val))
}
