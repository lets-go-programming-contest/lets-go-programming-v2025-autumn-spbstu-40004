package codingProcessor

import (
	"encoding/xml"
	"strconv"
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

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func (val ValCurs) Len() int {
	return len(val.Valutes)
}

func (val ValCurs) Swap(lhs, rhs int) {
	val.Valutes[lhs], val.Valutes[rhs] = val.Valutes[rhs], val.Valutes[lhs]
}

func (val ValCurs) Less(lhs, rhs int) bool {
	numI, errI := strconv.ParseFloat(val.Valutes[lhs].Value, 64)
	numJ, errJ := strconv.ParseFloat(val.Valutes[rhs].Value, 64)
	if errI != nil || errJ != nil {
		panic("Some errors in float parsing")
	}

	return numI > numJ
}
