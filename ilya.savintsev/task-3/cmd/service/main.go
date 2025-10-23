package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"

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
	ID        string   `xml:"ID,attr"`
	NumCode   int      `xml:"NumCode"`
	CharCode  string   `xml:"CharCode"`
	Nominal   int      `xml:"Nominal"`
	Name      string   `xml:"Name"`
	Value     string   `xml:"Value"`
	VunitRate string   `xml:"VunitRate"`
}

type ValuteShort struct {
	NumCode  int    `json:"num_code"`
	CharCode string `json:"char_code"`
	Value    string `json:"value"`
}

func main() {
	var fileDir string

	flag.StringVar(&fileDir, "config", "yaml", "Specifies the path to the config")
	flag.Parse()

	content, err := os.ReadFile(fileDir)
	if err != nil {
		fmt.Println("config does not exist")

		return
	}

	var config DirHandle

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		fmt.Println("unmarshal yaml error")

		return
	}

	valuteCurs, err := os.Open(config.InputFile)
	if err != nil {
		fmt.Println("opening xml error")
	}

	defer valuteCurs.Close()

	parser := xml.NewDecoder(valuteCurs)

	curs := new(ValCurs)

	for t, _ := parser.Token(); t != nil; t, _ = parser.Token() {
		if se, ok := t.(xml.StartElement); ok {
			if se.Name.Local == ValuteName {
				var item Valute
				parser.DecodeElement(&item, &se)
				if err = parser.DecodeElement(&item, &se); err != nil {
					fmt.Println("decode element error")

					return
				}
				curs.Valutes = append(curs.Valutes, item)
			}
		}
	}

	cursTemp := make([]ValuteShort, 0, len(curs.Valutes))

	for _, value := range curs.Valutes {
		valTemp := ValuteShort{
			NumCode:  value.NumCode,
			CharCode: value.CharCode,
			Value:    value.Value,
		}

		cursTemp = append(cursTemp, valTemp)
	}

	jsonData, err := json.MarshalIndent(cursTemp, "", "  ")
	if err != nil {
		fmt.Println("marshal json error")

		return
	}

	err = os.WriteFile(config.OutputFile, jsonData, 0600)
	if err != nil {
		fmt.Println("write file error")

		return
	}
}
