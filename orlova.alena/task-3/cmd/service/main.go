package main

import (
	"task-3/internal/config"
	"task-3/internal/converter"
)

func main() {
	cfg := config.Load()

	valCurs := converter.ParseXML(cfg.InputFile)

	currencies := converter.ConvertToCurrencies(valCurs)

	converter.WriteJSON(currencies, cfg.OutputFile)
}
