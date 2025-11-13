package currencyparser

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type DotFloat float64

type Currency struct {
	NumCode  int      `json:"num_code"  xml:"NumCode"`
	CharCode string   `json:"char_code" xml:"CharCode"`
	Value    DotFloat `json:"value"     xml:"Value"`
}

type Currencies struct {
	AllCurrencies []Currency `xml:"Valute"`
}

func Sort(currencies []Currency) []Currency {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies
}

func (dotFloat *DotFloat) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var str string

	err := decoder.DecodeElement(&str, &start)
	if err != nil {
		return fmt.Errorf("decoding xml: %w", err)
	}

	str = strings.ReplaceAll(str, ",", ".")

	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("parsing float '%s': %w", str, err)
	}

	*dotFloat = DotFloat(num)

	return nil
}
