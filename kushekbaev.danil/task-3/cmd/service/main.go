package main

import (
	"github.com/Z-1337/task-3/internal/cfgreader"
	"github.com/Z-1337/task-3/internal/currencymanager"
	"github.com/Z-1337/task-3/internal/currencyparser"
)

func main() {
	cfg, err := cfgreader.Parse()
	if err != nil {
		panic("Config parsing error: " + err.Error())
	}

	currencies, err := currencymanager.Read(cfg.Input)
	if err != nil {
		panic("Reading error: " + err.Error())
	}

	currencies.AllCurrencies = currencyparser.Sort(currencies.AllCurrencies)

	err = currencymanager.Write(cfg.Output, currencies)
	if err != nil {
		panic("Writing error: " + err.Error())
	}
}
