package main

import (
	"encoding/json"
	"flag"
	"sort"

	configHandler "github.com/belyaevEDU/task-3/internal/config_handling"
	ioHandler "github.com/belyaevEDU/task-3/internal/io_handling"
	xmlHandler "github.com/belyaevEDU/task-3/internal/xml_handling"
)

const (
	configFlagName         = "config"
	configFlagDesc         = "Configuration file path"
	configFlagDefaultValue = ""
	configFlagPanicMessage = "config file path via flag not specified"
)

func main() {
	var configFilePath string

	flag.StringVar(&configFilePath, configFlagName, configFlagDefaultValue, configFlagDesc)
	flag.Parse()

	flagExists := false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == configFlagName {
			flagExists = true
		}
	})

	if !flagExists {
		panic(configFlagPanicMessage)
	}

	config, err := configHandler.LoadConfig(configFilePath)
	if err != nil {
		panic(err.Error())
	}

	currencies, err := xmlHandler.ParseXML(config.InputFile)
	if err != nil {
		panic(err.Error())
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	jsonMarshalled, err := json.MarshalIndent(currencies, "", "\t")
	if err != nil {
		panic(err.Error())
	}

	err = ioHandler.WriteStringToFile(config.OutputFile, jsonMarshalled)
	if err != nil {
		panic(err.Error())
	}
}
