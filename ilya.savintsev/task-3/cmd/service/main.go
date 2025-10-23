package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
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

var errOpening = errors.New("opening xml file error")

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

func parseXML(filepath string) (*ValCurs, error) {
	valuteCurs, err := os.Open(filepath)
	if err != nil {
		return nil, errOpening
	}

	defer func() {
		if err = valuteCurs.Close(); err != nil {
			return
		}
	}()

	parser := xml.NewDecoder(valuteCurs)

	curs := new(ValCurs)

	for t, _ := parser.Token(); t != nil; t, _ = parser.Token() {
		if se, ok := t.(xml.StartElement); ok {
			if se.Name.Local == ValuteName {
				var item Valute

				if err = parser.DecodeElement(&item, &se); err != nil {
					fmt.Println("decode element error")

					return nil, errOpening
				}

				curs.Valutes = append(curs.Valutes, item)
			}
		}
	}

	return curs, nil
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

	curs, err := parseXML(config.InputFile)
	if err != nil {
		fmt.Println(err)

		return
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
