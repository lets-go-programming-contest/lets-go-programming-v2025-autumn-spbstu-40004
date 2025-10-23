package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v2"
)

const (
	ownerReadWrite = 0o600
	allReadWrite   = 0o755
	ValuteName     = "Valute"
)

type DirHandle struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	XMLName   xml.Name `xml:"Valute"`
	ID        string   `xml:"ID,attr"`
	NumCode   string   `xml:"NumCode"`
	CharCode  string   `xml:"CharCode"`
	Nominal   int      `xml:"Nominal"`
	Name      string   `xml:"Name"`
	Value     string   `xml:"Value"`
	VunitRate string   `xml:"VunitRate"`
}

type ValuteShort struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

var (
	errOpenXML  = errors.New("open xml error")
	errDecdXML  = errors.New("decode xml error")
	errMarhJSON = errors.New("marshal json error")
)

func parseXML(filepath string) (*ValCurs, error) {
	valuteCurs, err := os.Open(filepath)
	if err != nil {
		return nil, errOpenXML
	}

	defer func() {
		if err = valuteCurs.Close(); err != nil {
			panic(err)
		}
	}()

	decoder := xml.NewDecoder(valuteCurs)
	decoder.CharsetReader = charset.NewReaderLabel

	var curs ValCurs
	if err := decoder.Decode(&curs); err != nil {
		return nil, errDecdXML
	}

	return &curs, nil
}

func createJSON(curs *ValCurs) ([]byte, error) {
	cursTemp := make([]ValuteShort, 0, len(curs.Valutes))

	for _, value := range curs.Valutes {
		numCode, err := strconv.Atoi(value.NumCode)
		if err != nil {
			return nil, fmt.Errorf("json int error: %w", err)
		}

		valueWithDot := strings.ReplaceAll(value.Value, ",", ".")

		floatValue, err := strconv.ParseFloat(valueWithDot, 64)
		if err != nil {
			return nil, fmt.Errorf("parse value error: %w", err)
		}

		valTemp := ValuteShort{
			NumCode:  numCode,
			CharCode: value.CharCode,
			Value:    floatValue,
		}

		cursTemp = append(cursTemp, valTemp)
	}

	sort.Slice(cursTemp, func(i, j int) bool {
		return cursTemp[i].Value > cursTemp[j].Value
	})

	jsonData, err := json.MarshalIndent(cursTemp, "", "  ")
	if err != nil {
		return nil, errMarhJSON
	}

	return jsonData, nil
}

func saveToFile(jsonData []byte, outputFile string) error {
	dir := filepath.Dir(outputFile)

	if err := os.MkdirAll(dir, allReadWrite); err != nil {
		return fmt.Errorf("create directory error: %w", err)
	}

	err := os.WriteFile(outputFile, jsonData, ownerReadWrite)
	if err != nil {
		return fmt.Errorf("write file error: %w", err)
	}

	return nil
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

	jsonData, err := createJSON(curs)
	if err != nil {
		fmt.Println(err)

		return
	}

	err = saveToFile(jsonData, config.OutputFile)
	if err != nil {
		fmt.Println("write file error")
		return
	}
}
