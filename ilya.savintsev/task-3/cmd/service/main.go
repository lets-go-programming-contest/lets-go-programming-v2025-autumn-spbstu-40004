package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"io"
	"os"
	"path/filepath"
	"sort"

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
	XMLName   xml.Name `json:"-"         xml:"Valute"`
	ID        string   `json:"-"         xml:"ID,attr"`
	NumCode   int      `json:"num_code"  xml:"NumCode"`
	CharCode  string   `json:"char_code" xml:"CharCode"`
	Nominal   int      `json:"-"         xml:"Nominal"`
	Name      string   `json:"-"         xml:"Name"`
	Value     float64  `json:"value"     xml:"Value"`
	VunitRate string   `json:"-"         xml:"VunitRate"`
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	errOpenXML = errors.New("no such file or directory")
	errReadXML = errors.New("invalid file")
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

	content, err := io.ReadAll(valuteCurs)
	if err != nil {
		return nil, errReadXML
	}

	contentWithDots := bytes.ReplaceAll(content, []byte(","), []byte("."))

	decoder := xml.NewDecoder(bytes.NewReader(contentWithDots))
	decoder.CharsetReader = charset.NewReaderLabel

	var curs ValCurs
	if err := decoder.Decode(&curs); err != nil {
		return nil, errDecdXML
	}

	return &curs, nil
}

var errMarsJSON = errors.New("cant marshall json")

func createJSON(curs *ValCurs) ([]byte, error) {
	cursTemp := make([]Valute, 0, len(curs.Valutes))

	for _, value := range curs.Valutes {
		valTemp := Valute{
			XMLName:   xml.Name{Local: "Valute", Space: ""},
			ID:        value.ID,
			NumCode:   value.NumCode,
			CharCode:  value.CharCode,
			Nominal:   value.Nominal,
			Name:      value.Name,
			Value:     value.Value,
			VunitRate: value.VunitRate,
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
