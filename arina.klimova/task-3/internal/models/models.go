package models

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type currencyXML struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	ValueStr string `xml:"Value"`
}

type Currency struct {
	NumCode  string  `xml:"-"`
	CharCode string  `xml:"-"`
	Value    float64 `xml:"-"`
}

func (c *Currency) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var temp currencyXML
	if err := decoder.DecodeElement(&temp, &start); err != nil {
		return err
	}

	c.NumCode = temp.NumCode
	c.CharCode = temp.CharCode

	valueStr := strings.Replace(temp.ValueStr, ",", ".", -1)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return err
	}
	c.Value = value

	return nil
}

type ValCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

type CurrencyOutput struct {
	NumCode  string  `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}
