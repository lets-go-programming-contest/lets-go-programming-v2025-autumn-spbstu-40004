package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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

type ValuteJSON struct {
	NumCode  int64   `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type ErrInvalidNumCode struct {
	NumCode string
	Valute  Valute
}

func (e ErrInvalidNumCode) Error() string {
	return fmt.Sprintf("invalid NumCode '%s' for element: %+v", e.NumCode, e.Valute)
}

func ParseValuteXML(path string) (ValCurs, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return ValCurs{}, err
	}

	defer func() {
		if err := xmlFile.Close(); err != nil {
			panic(fmt.Sprintf("Error closing XML file: %v\n", err))
		}
	}()

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

func ConvertValutesToJSON(valutes []Valute) ([]ValuteJSON, error) {
	valutesJson := make([]ValuteJSON, 0, len(valutes))
	for _, valute := range valutes {
		value, err := parseValue(valute.Value)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Parsing %+v\n", valute)
		var numCode int64
		if valute.NumCode == "" {
			numCode = 0
		} else {
			var err error
			numCode, err = strconv.ParseInt(valute.NumCode, 10, 64)
			if err != nil {
				return nil, ErrInvalidNumCode{NumCode: valute.NumCode, Valute: valute}
			}
		}
		valutesJson = append(valutesJson, ValuteJSON{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	return valutesJson, nil
}

func SaveToJSON(valutesJson []ValuteJSON, outputPath string) error {

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}
	jsonFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			panic(fmt.Sprintf("Error closing JSON file: %v\n", err))
		}
	}()

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
