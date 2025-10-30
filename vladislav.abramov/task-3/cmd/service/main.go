package main

import (
	"flag"

	"currency-processor/internal/parser"
	"currency-processor/internal/processor"
	"currency-processor/internal/ioutils"
)

func main() {
	configPath := flag.String("config", "", "Path to YAML configuration file")
	flag.Parse()

	if *configPath == "" {
		panic("configuration file path must be provided using --config flag")
	}

	cfg, err := config.LoadConfiguration(*configPath)
	if err != nil {
		panic("failed to load configuration: " + err.Error())
	}

	currencyData, err := parser.ParseCurrencyData(cfg.InputFile)
	if err != nil {
		panic("failed to parse currency data: " + err.Error())
	}

	sortedCurrencies, err := processor.SortCurrenciesByValue(currencyData.Valutes)
	if err != nil {
		panic("failed to process currency data: " + err.Error())
	}

	err = writer.WriteJSONOutput(sortedCurrencies, cfg.OutputFile)
	if err != nil {
		panic("failed to write output file: " + err.Error())
	}
}
