package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v2"
)

type DirHandle struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute
}

const ValuteName = "Valute"

type Valute struct {
	XMLName   xml.Name `xml:"Valute"`
	Id        []rune   `xml:"ID,attr"`
	NumCode   int      `xml:"NumCode"`
	CharCode  string   `xml:"CharCode"`
	Nominal   int      `xml:"Nominal"`
	Name      string   `xml:"Name"`
	Value     string  `xml:"Value"`
	VunitRate string  `xml:"VunitRate"`
}

func main() {
	var fileDir string
	flag.StringVar(&fileDir, "config", "yaml", "Specifies the path to the config")
	flag.Parse()

	fmt.Println(fileDir)

	content, err := os.ReadFile(fileDir)
	if err != nil {
		fmt.Println("config does not exist")

		return
	}

	var y DirHandle
	err = yaml.Unmarshal(content, &y)
	if err != nil {
		fmt.Printf("error: %v", err)

		return
	}

	valuteCurs, err := os.Open(y.InputFile)
	if err != nil {
		fmt.Println(err)
	}

	defer valuteCurs.Close()

	parser := xml.NewDecoder(valuteCurs)
	parser.CharsetReader = charset.NewReaderLabel

	curs := new(ValCurs)

	for t, _ := parser.Token(); t != nil; t, _ = parser.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == ValuteName {
				var item Valute
				parser.DecodeElement(&item, &se)
				//fmt.Println("Id: " + item.Name)
				curs.Valutes = append(curs.Valutes, item)
				fmt.Println(item.Id)
			}
		}
	}
}
