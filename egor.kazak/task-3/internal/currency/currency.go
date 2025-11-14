package currency

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func (c *Currency) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var aux struct {
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		ValueStr string `xml:"Value"`
	}

	err := decoder.DecodeElement(&aux, &start)
	if err != nil {
		return fmt.Errorf("failed to decode XML element: %w", err)
	}

	value, err := strconv.ParseFloat(strings.ReplaceAll(aux.ValueStr, ",", "."), 64)
	if err != nil {
		return fmt.Errorf("failed to parse float from %q: %w", aux.ValueStr, err)
	}

	*c = Currency{
		NumCode:  aux.NumCode,
		CharCode: aux.CharCode,
		Value:    value,
	}

	return nil
}

type ValCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

func ParseXML(data []byte) ([]Currency, error) {
	reader := strings.NewReader(string(data))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs

	err := decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ValCurs: %w", err)
	}

	sort.Slice(valCurs.Currencies, func(i, j int) bool {
		return valCurs.Currencies[i].Value > valCurs.Currencies[j].Value
	})

	return valCurs.Currencies, nil
}
