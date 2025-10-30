package main

import (
	"flag"
	"os"

	"github.com/ZakirovMS/task-3/internal/codingProcessor"
	"gopkg.in/yaml.v3"
)

func main() {
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

	var inData codingProcessor.Currency
	err = yaml.Unmarshal(inFile, &inData)
	if err != nil {
		panic("Some errors in decoding input file")
	}

}
