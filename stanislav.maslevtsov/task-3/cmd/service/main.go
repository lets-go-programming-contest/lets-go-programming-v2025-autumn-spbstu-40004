package main

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/jambii1/task-3/internal/config"
	"github.com/jambii1/task-3/internal/currency"
)

var errInvalidCommandParameters = errors.New("invalid command parameters")

func main() {
	if len(os.Args) != 2 || os.Args[0] != "-config" {
		fmt.Println(errInvalidCommandParameters)
	}

	config, err := config.Parse(os.Args[1])
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
