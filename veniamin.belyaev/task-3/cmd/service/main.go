package main

import (
	"encoding/json"
	"flag"
	"strings"

	configHandler "github.com/belyaevEDU/task-3/internal/config_handling"
	ioHandler "github.com/belyaevEDU/task-3/internal/io_handling"
	xmlHandler "github.com/belyaevEDU/task-3/internal/xml_handling"
)

func main() {
	var configFilePath string

	flag.StringVar(&configFilePath, "config", "none", "Configuration file path")
	flag.Parse()

	if strings.Compare(configFilePath, "") == 0 {
		panic("config file path via flag not specified")
	}

	config, err := configHandler.LoadConfig(configFilePath)
	if err != nil {
		panic(err.Error())
	}

	currenciesXML, err := xmlHandler.ParseXML(config.InputFile)
	if err != nil {
		panic(err.Error())
	}

	currenciesJSON, err := xmlHandler.ConvertXMLStructsToJSON(*currenciesXML)
	if err != nil {
		panic(err.Error())
	}

	jsonMarshalled, err := json.MarshalIndent(currenciesJSON, "", "\t")
	if err != nil {
		panic(err.Error())
	}

	err = ioHandler.WriteStringToFile(config.OutputFile, jsonMarshalled)
	if err != nil {
		panic(err.Error())
	}
}
