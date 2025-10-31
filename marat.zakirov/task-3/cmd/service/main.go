package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"

	"github.com/ZakirovMS/task-3/internal/codingprocessor"
	"github.com/ZakirovMS/task-3/internal/currencyprocessor"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

func main() {
	var nFlag = flag.String("config", "", "Path to YAML config file")
	const permission = 0o666

	flag.Parse()
	configFile, err := os.ReadFile(*nFlag)
	if err != nil {
		panic("Some errors in getting config file")
	}

	var ioPath codingprocessor.PathHolder
	err = yaml.Unmarshal(configFile, &ioPath)
	if err != nil || ioPath.InPath == "" {
		panic("Some errors in decoding config file")
	}

	inFile, err := os.Open(ioPath.InPath)
	if err != nil {
		panic("Some errors in reading YAML input file")
	}

	decoder := xml.NewDecoder(inFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var inData currencyprocessor.ValCurs
	err = decoder.Decode(&inData)
	if err != nil {
		panic("Some errors in decoding YAML input file")
	}

	currencyprocessor.SortValue(&inData)
	jsonData := codingprocessor.ConvertXMLToJSON(inData)

	outData, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		panic("Some errors in json encoding")
	}

	err = os.MkdirAll(filepath.Dir(ioPath.OutPath), 0o755)
	if err != nil {
		panic("Some errors in creating directories")
	}

	err = os.WriteFile(ioPath.OutPath, outData, permission)
	if err != nil {
		panic("Some errors in file writing")
	}
}
