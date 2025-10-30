package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"

	"github.com/ZakirovMS/task-3/internal/codingProcessor"
	"github.com/ZakirovMS/task-3/internal/currencyProcessor"
	"gopkg.in/yaml.v3"
)

func main() {
	const permisions = 0o666
	var nFlag = flag.String("config", "", "Path to YAML config file")

	flag.Parse()
	configFile, err := os.ReadFile(*nFlag)
	if err != nil {
		panic("Some errors in getting config file")
	}

	var ioPath codingProcessor.PathHolder
	err = yaml.Unmarshal(configFile, &ioPath)
	if err != nil || ioPath.InPath == "" {
		panic("Some errors in decoding config file")
	}

	inFile, err := os.ReadFile(ioPath.InPath)
	if err != nil {
		panic("Some errors in reading YAML input file")
	}

	var inData currencyProcessor.ValCurs
	err = xml.Unmarshal(inFile, &inData)
	if err != nil {
		panic("Some errors in decoding input file")
	}

	currencyProcessor.SortValue(&inData)
	jsonData := codingProcessor.ConvertXmlToJson(inData)

	outData, err := json.Marshal(jsonData)
	if err != nil {
		panic("Some errors in json encoding")
	}

	err = os.WriteFile(ioPath.OutPath, outData, permisions)
	if err != nil {
		panic("Some errors in file writing")
	}
}
