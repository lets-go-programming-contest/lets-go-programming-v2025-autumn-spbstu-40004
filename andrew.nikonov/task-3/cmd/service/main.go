package main

import (
	"flag"

	"github.com/ysffmn/task-3/internal/config"
	"github.com/ysffmn/task-3/internal/parser"
	"github.com/ysffmn/task-3/internal/processor"
	"github.com/ysffmn/task-3/internal/writer"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("Config flag is required")
	}

	cfg := config.Load(*configPath)

	valCurs := parser.ParseXMLFile(cfg.InputFile)
	currencies := parser.ConvertToCurrencies(valCurs)
	sortedCurrencies := processor.SortCurrenciesDesc(currencies)
	writer.WriteJSONToFile(sortedCurrencies, cfg.OutputFile)
}
