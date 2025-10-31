package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type currencyXML struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	ValueStr string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `xml:"-"`
	CharCode string  `xml:"-"`
	Value    float64 `xml:"-"`
}

func (c *Currency) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var temp currencyXML
	if err := decoder.DecodeElement(&temp, &start); err != nil {
		return fmt.Errorf("decode XML element: %w", err)
	}

	if temp.NumCode != "" {
		numCode, err := strconv.Atoi(temp.NumCode)
		if err != nil {
			return fmt.Errorf("convert num code: %w", err)
		}

		c.NumCode = numCode
	} else {
		c.NumCode = 0
	}

	if temp.CharCode != "" {
		c.CharCode = temp.CharCode
	} else {
		c.CharCode = ""
	}

	valueStr := strings.ReplaceAll(temp.ValueStr, ",", ".")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("parse value: %w", err)
	}

	c.Value = value

	return nil
}

type ValCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

type CurrencyOutput struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}
