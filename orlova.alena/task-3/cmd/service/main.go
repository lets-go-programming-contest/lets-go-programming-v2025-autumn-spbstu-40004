package main

import (
	"task-3/internal/config"
	"task-3/internal/converter"
)

func main() {
	cfg := config.Load()

	valCurs, err := converter.ParseXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currencies := converter.ConvertToCurrencies(valCurs)
	if err != nil {
		panic(err)
	}

	err = converter.WriteJSON(currencies, cfg.OutputFile)
	if err != nil {
		panic(err)
	}
}
