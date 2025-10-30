package models

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type Currency struct {
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	ValueStr string  `xml:"Value"`
	Value    float64 `json:"value"`
}

func (c *Currency) ConvertFloatValue() error {
	valueStr := strings.Replace(c.ValueStr, ",", ".", -1)
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
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}
