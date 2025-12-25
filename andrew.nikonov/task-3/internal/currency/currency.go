package currency

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string  `xml:"ID,attr"`
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Nominal  int     `xml:"Nominal"`
	Name     string  `xml:"Name"`
	Value    float64 `json:"value"     xml:"Value"`
}

func (v *Valute) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	type valuteAlias struct {
		ID       string `xml:"ID,attr"`
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Nominal  int    `xml:"Nominal"`
		Name     string `xml:"Name"`
		Value    string `xml:"Value"`
	}

	var alias valuteAlias

	err := decoder.DecodeElement(&alias, &start)
	if err != nil {
		return fmt.Errorf("failed to decode XML element: %w", err)
	}

	v.ID = alias.ID
	v.NumCode = alias.NumCode
	v.CharCode = alias.CharCode
	v.Nominal = alias.Nominal
	v.Name = alias.Name

	normalized := strings.Replace(alias.Value, ",", ".", 1)

	parsedValue, parseErr := strconv.ParseFloat(normalized, 64)
	if parseErr != nil {
		return fmt.Errorf("failed to parse value: %w", parseErr)
	}

	v.Value = parsedValue

	return nil
}
