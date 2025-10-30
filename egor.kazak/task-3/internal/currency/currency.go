package currency

import (
	"encoding/xml"
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

func (c *Currency) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var aux struct {
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		ValueStr string `xml:"Value"`
	}
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	v := strings.Replace(aux.ValueStr, ",", ".", -1)
	value, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return err
	}
	*c = Currency{
		NumCode:  aux.NumCode,
		CharCode: aux.CharCode,
		Value:    value,
	}
	return nil
}

type ValCurs struct {
	XMLName    xml.Name    `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

func ParseXML(data []byte) ([]Currency, error) {
	reader := strings.NewReader(string(data))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel // ← ключевая строка

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, err
	}

	sort.Slice(valCurs.Currencies, func(i, j int) bool {
		return valCurs.Currencies[i].Value > valCurs.Currencies[j].Value
	})

	return valCurs.Currencies, nil
}
