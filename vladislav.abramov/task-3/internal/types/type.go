package types

import "encoding/xml"

type CurrencyData struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	XMLName  xml.Name `json:"-"         xml:"Valute"`
	ID       string   `json:"-"         xml:"ID,attr"`
	NumCode  int      `json:"num_code"  xml:"NumCode"`
	CharCode string   `json:"char_code" xml:"CharCode"`
	Nominal  int      `json:"-"         xml:"Nominal"`
	Name     string   `json:"-"         xml:"Name"`
	Value    float64  `json:"value"     xml:"Value"`
}

type CurrencyOutput struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}
