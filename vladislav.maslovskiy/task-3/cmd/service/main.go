package main

import (
	"encoding/json"
	"flag"
	"sort"

	configHandler "github.com/vladislavmaslovskiy/task-3/internal/config_handling"
	ioHandler "github.com/vladislavmaslovskiy/task-3/internal/io_handling"
	xmlHandler "github.com/vladislavmaslovskiy/task-3/internal/xml_handling"
)

const (
	configFlagName         = "config"
	configFlagDesc         = "Configuration file path"
	configFlagDefaultValue = ""
	configFlagPanicMessage = "config file path via flag not specified"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, configFlagName, configFlagDefaultValue, configFlagDesc)
	flag.Parse()

	flagProvided := false

	flag.Visit(func(currentFlag *flag.Flag) {
		if currentFlag.Name == configFlagName {
			flagProvided = true
		}
	})

	if !flagProvided {
		panic(configFlagPanicMessage)
	}

	settings, err := configHandler.LoadConfiguration(configPath)
	if err != nil {
		panic(err.Error())
	}

	currencyData, err := xmlHandler.ParseXML(settings.InputFile)
	if err != nil {
		panic(err.Error())
	}

	sort.Slice(currencyData, func(firstIndex, secondIndex int) bool {
		return currencyData[firstIndex].Value > currencyData[secondIndex].Value
	})

	jsonData, err := json.MarshalIndent(currencyData, "", "\t")
	if err != nil {
		panic(err.Error())
	}

	err = ioHandler.WriteDataToFile(settings.OutputFile, jsonData)
	if err != nil {
		panic(err.Error())
	}
}
