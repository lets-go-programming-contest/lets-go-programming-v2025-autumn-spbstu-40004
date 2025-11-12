package main

import (
	"flag"

	"github.com/lolnyok/task-3/internal/confdecoder"
	"github.com/lolnyok/task-3/internal/indecoder"
	"github.com/lolnyok/task-3/internal/outcoder"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	config, err := confdecoder.ParseConfiguration(*configPath)
	if err != nil {
		panic(err.Error())
	}

	currencyData, err := indecoder.ProcessCurrencyFile(config.SourcePath)
	if err != nil {
		panic(err.Error())
	}

	err = outcoder.SaveCurrencyData(config.TargetPath, currencyData)
	if err != nil {
		panic(err.Error())
	}
}
