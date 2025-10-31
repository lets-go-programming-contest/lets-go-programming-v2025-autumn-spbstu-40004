package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID        string `xml:"ID,attr"`
	NumCode   string `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   string `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

type ValuteJson struct {
	NumCode  int64   `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ParseValuteXML(path string) (ValCurs, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return ValCurs{}, err
	}
	defer xmlFile.Close()

	var valCurs ValCurs
	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		if encoding == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return input, nil
	}
	if err := decoder.Decode(&valCurs); err != nil {
		return ValCurs{}, err
	}

	return valCurs, nil
}

func convertValutesToJson(valutes []Valute) ([]ValuteJson, error) {
	var valutesJson []ValuteJson
	for _, v := range valutes {
		value, err := parseValue(v.Value)
		if err != nil {
			return nil, err
		}

		var numCode int64
		if v.NumCode == "" {
			numCode = 0
		} else {
			var err error
			numCode, err = strconv.ParseInt(v.NumCode, 10, 64)
			if err != nil {
				return nil, err
			}
		}
		valutesJson = append(valutesJson, ValuteJson{
			NumCode:  numCode,
			CharCode: v.CharCode,
			Value:    value,
		})
	}
	return valutesJson, nil
}

func SaveToJson(valutes []Valute, outputPath string) error {
	valutesJson, err := convertValutesToJson(valutes)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}
	jsonFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(valutesJson); err != nil {
		return err
	}

	return nil
}

func parseValue(s string) (float64, error) {
	s = strings.Replace(s, ",", ".", 1)
	return strconv.ParseFloat(s, 64)
}
