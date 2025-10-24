package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
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

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	errOpenXML = errors.New("no such file or directory")
	errDecdXML = errors.New("invalid encoding")
)

func parseXML(filepath string) (*ValCurs, error) {
	valuteCurs, err := os.Open(filepath)
	if err != nil {
		return nil, errOpenXML
	}

	defer func() {
		panicIfErr(valuteCurs.Close())
	}()

	decoder := xml.NewDecoder(valuteCurs)
	decoder.CharsetReader = charset.NewReaderLabel

	var curs ValCurs
	if err := decoder.Decode(&curs); err != nil {
		return nil, errDecdXML
	}

	return &curs, nil
}

var (
	errMarsJSON = errors.New("cant marshall json")
	errConvJSON = errors.New("cant convert numbers")
)

func createJSON(curs *ValCurs) ([]byte, error) {
	cursTemp := make([]ValuteShort, 0, len(curs.Valutes))

	for _, value := range curs.Valutes {
		numCode, err := strconv.Atoi(value.NumCode)
		if err != nil {
			return nil, errConvJSON
		}

		valueWithDot := strings.ReplaceAll(value.Value, ",", ".")

		floatValue, err := strconv.ParseFloat(valueWithDot, 64)
		if err != nil {
			return nil, errConvJSON
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
		return nil, errMarsJSON
	}

	return jsonData, nil
}

var (
	errDirSave = errors.New("unable create directory")
	errWrtSave = errors.New("unable save file")
)

func saveToFile(jsonData []byte, outputFile string) error {
	dir := filepath.Dir(outputFile)

	if err := os.MkdirAll(dir, allReadWrite); err != nil {
		return errDirSave
	}

	err := os.WriteFile(outputFile, jsonData, ownerReadWrite)
	if err != nil {
		return errWrtSave
	}

	return nil
}

func main() {
	var fileDir string

	flag.StringVar(&fileDir, "config", "yaml", "Specifies the path to the config")
	flag.Parse()

	content, err := os.ReadFile(fileDir)
	if err != nil {
		panic("no such file or directory")
	}

	var config DirHandle

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic("did not find expected key")
	}

	curs, err := parseXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	jsonData, err := createJSON(curs)
	if err != nil {
		panic(err)
	}

	err = saveToFile(jsonData, config.OutputFile)
	if err != nil {
		panic(err)
	}
}
