package main

import (
	"github.com/arinaklimova/task-3/internal/config"
	"github.com/arinaklimova/task-3/internal/parser"
	"github.com/arinaklimova/task-3/internal/sorter"
	"github.com/arinaklimova/task-3/internal/writer"
)

func main() {
	cfg := config.LoadConfig()

	valCurs, err := parser.ParseXML(cfg.InputFile)
	if err != nil {
		panic("failed to parse XML file: " + err.Error())
	}

	sorter.SortCurrenciesByValueDesc(valCurs.Currencies)

	if err := writer.WriteJSON(valCurs.Currencies, cfg.OutputFile); err != nil {
		panic("failed to write JSON file: " + err.Error())
	}
}
