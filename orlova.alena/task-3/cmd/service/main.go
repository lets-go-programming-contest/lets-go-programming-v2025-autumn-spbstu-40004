package main

import (
	"task-3/internal/config"
	"task-3/internal/converter"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	valCurs, err := converter.ParseXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currencies := converter.ConvertToCurrencies(valCurs)

	err = converter.WriteJSON(currencies, cfg.OutputFile)
	if err != nil {
		panic(err)
	}
}
