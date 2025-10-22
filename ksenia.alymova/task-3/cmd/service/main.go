package main

import (
	"flag"

	"github.com/Ksenia-rgb/task-3/internal/confdecoder"
	"github.com/Ksenia-rgb/task-3/internal/indecoder"
	"github.com/Ksenia-rgb/task-3/internal/outcoder"
)

const (
	configFlag    string = "config"
	configDefault string = "resources/config/config.yaml"
	configHelp    string = "direction config file for task3"
)

func main() {
	flagConfig := flag.String(configFlag, configDefault, configHelp)
	flag.Parse()

	config, err := confdecoder.ConfigProcess(flagConfig)
	if err != nil {
		panic(err)
	}

	inputData, err := indecoder.InputProcess(config.InputFile)
	if err != nil {
		panic(err)
	}

	err = outcoder.OutputProcess(config.OutputFile, inputData)
	if err != nil {
		panic(err)
	}
}
