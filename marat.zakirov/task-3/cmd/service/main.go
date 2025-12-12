package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"

	"github.com/ZakirovMS/task-3/internal/currencyprocessor"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

func main() {
	nFlag := flag.String("config", "resources/config/config.yaml", "Path to YAML config file")

	const dirPerm, filePerm = 0o755, 0o666

	flag.Parse()

	configFile, err := os.ReadFile(*nFlag)
	if err != nil {
		panic("Some errors in getting config file")
	}

	var ioPath currencyprocessor.PathHolder

	err = yaml.Unmarshal(configFile, &ioPath)
	if err != nil || ioPath.InPath == "" {
		panic("Some errors in decoding config file, like did not find expected key")
	}

	inFile, err := os.Open(ioPath.InPath)
	if err != nil {
		panic("Some errors in reading YAML input file, like no such file or directory")
	}

	decoder := xml.NewDecoder(inFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var inData currencyprocessor.ValCurs

	err = decoder.Decode(&inData)
	if err != nil {
		panic("Some errors in decoding YAML input file")
	}

	currencyprocessor.SortValue(&inData)

	outData, err := json.MarshalIndent(inData.Valutes, "", "  ")
	if err != nil {
		panic("Some errors in json encoding")
	}

	err = os.MkdirAll(filepath.Dir(ioPath.OutPath), dirPerm)
	if err != nil {
		panic("Some errors in creating directories")
	}

	err = os.WriteFile(ioPath.OutPath, outData, filePerm)
	if err != nil {
		panic("Some errors in file writing")
	}
}
