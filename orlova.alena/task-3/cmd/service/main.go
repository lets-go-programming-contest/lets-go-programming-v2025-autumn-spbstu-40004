package main

import (
	"task-3/internal/config"
	"task-3/internal/converter"
)

func main() {
	cfg := config.Load()

	valCurs := converter.ParseXML(cfg.InputFile)

	currencies := converter.ConvertToCurrencies(valCurs)

	err := converter.WriteJSON(currencies, cfg.OutputFile)
	if err != nil {
		panic(err)
	}
}
