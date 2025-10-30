package models

import "encoding/xml"

type Currency struct {
	NumCode  string  `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Nominal  int     `xml:"Nominal"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
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
