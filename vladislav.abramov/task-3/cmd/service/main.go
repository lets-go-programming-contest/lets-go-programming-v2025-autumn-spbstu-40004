package main

import (
	"flag"

	"github.com/15446-rus75/task-3/internal/ioutils"
	"github.com/15446-rus75/task-3/internal/parser"
	"github.com/15446-rus75/task-3/internal/processor"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to YAML configuration file")
	flag.Parse()

	if *configPath == "" {
		panic("configuration file path must be provided using --config flag")
	}

	cfg, err := ioutils.LoadConfiguration(*configPath)
	if err != nil {
		panic("failed to load configuration: " + err.Error())
	}

	currencyData, err := parser.ParseCurrencyData(cfg.InputFile)
	if err != nil {
		panic("failed to parse currency data: " + err.Error())
	}

	sortedCurrencies := processor.SortCurrenciesByValue(currencyData.Valutes)

	err = ioutils.WriteJSONOutput(sortedCurrencies, cfg.OutputFile)
	if err != nil {
		panic("failed to write output file: " + err.Error())
	}
}
