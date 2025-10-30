package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"

	"github.com/ZakirovMS/task-3/internal/codingprocessor"
	"github.com/ZakirovMS/task-3/internal/currencyprocessor"
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

	inFile, err := os.ReadFile(ioPath.InPath)
	if err != nil {
		panic("Some errors in reading YAML input file")
	}

	var inData currencyprocessor.ValCurs
	err = xml.Unmarshal(inFile, &inData)
	if err != nil {
		panic("Some errors in decoding input file")
	}

	currencyprocessor.SortValue(&inData)
	jsonData := codingprocessor.ConvertXMLToJSON(inData)

	outData, err := json.Marshal(jsonData)
	if err != nil {
		panic("Some errors in json encoding")
	}

	err = os.WriteFile(ioPath.OutPath, outData, permission)
	if err != nil {
		panic("Some errors in file writing")
	}
}
