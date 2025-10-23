package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	errReadXML  = errors.New("read xml error")
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

	content, err := io.ReadAll(valuteCurs)
	if err != nil {
		return nil, errReadXML
	}

	transformed := strings.ReplaceAll(string(content), ",", ".")
	reader := strings.NewReader(transformed)

	decoder := xml.NewDecoder(reader)
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

		floatValue, err := strconv.ParseFloat(value.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("json float error: %w", err)
		}

		valTemp := ValuteShort{
			NumCode:  numCode,
			CharCode: value.CharCode,
			Value:    floatValue,
		}

		cursTemp = append(cursTemp, valTemp)
	}

	jsonData, err := json.MarshalIndent(cursTemp, "", "  ")
	if err != nil {
		return nil, errMarhJSON
	}

	return jsonData, nil
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

	dir := filepath.Dir(config.OutputFile)

	err = os.MkdirAll(dir, allReadWrite)
	if err != nil {
		fmt.Println("write file error")

		return
	}

	err = os.WriteFile(config.OutputFile, jsonData, ownerReadWrite)
	if err != nil {
		fmt.Println("write file error")

		return
	}
}
