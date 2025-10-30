package types

import "encoding/xml"

type CurrencyData struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	XMLName  xml.Name `xml:"Valute"   json:"-"`
	ID       string   `xml:"ID,attr"  json:"-"`
	NumCode  int      `xml:"NumCode"  json:"num_code"`
	CharCode string   `xml:"CharCode" json:"char_code"`
	Nominal  int      `xml:"Nominal"  json:"-"`
	Name     string   `xml:"Name"     json:"-"`
	Value    float64  `xml:"Value"    json:"value"`
}

type CurrencyOutput struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}
