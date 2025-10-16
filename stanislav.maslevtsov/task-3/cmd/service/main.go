package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/jambii1/task-3/internal/configparser"
	"github.com/jambii1/task-3/internal/currency"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "-config" {
		fmt.Println("invalid command parameters")
	}

	config, err := configparser.Parse(os.Args[2])
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
