package main

import (
	"flag"
	"slices"

	"github.com/jambii1/task-3/internal/config"
	"github.com/jambii1/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "./config.yaml", "config path")
	flag.Parse()

	config, err := config.Parse(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.Parse(config.InputFile)
	if err != nil {
		panic(err)
	}

	slices.SortFunc(currencies.Data, currency.Compare)

	err = currency.Write(config.OutputFile, currencies)
	if err != nil {
		panic(err)
	}
}
