package main

import (
	"errors"
	"flag"
	"fmt"
	"slices"

	"github.com/jambii1/task-3/internal/config"
	"github.com/jambii1/task-3/internal/currency"
)

var errInvalidCommandParameters = errors.New("invalid command parameters")

func main() {
	configPath := flag.String("config", "", "config path")
	flag.Parse()

	if *configPath == "" {
		fmt.Println(errInvalidCommandParameters)

		return
	}

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
