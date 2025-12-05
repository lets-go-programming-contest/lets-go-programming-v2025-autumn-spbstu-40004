package models

import "encoding/xml"

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	XMLName xml.Name `json:"-"         xml:"Valute"`
	Name    string   `json:"-"         xml:"Name"`
	Value   string   `json:"value"     xml:"Value"`
}

type Currency struct {
	Value float64 `json:"value"`
}
