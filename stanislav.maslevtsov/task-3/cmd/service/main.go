package main

import (
	"fmt"
	"os"

	"github.com/jambii1/task-3/internal/configparser"
	curproc "github.com/jambii1/task-3/internal/currenciesprocessing"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "-config" {
		fmt.Println("invalid command parameters")
	}

	config, err := configparser.ParseConfig(os.Args[2])
	if err != nil {
		panic(err)
	}

	currencies, err := curproc.ParseCurrencies(config.InputFile)
	if err != nil {
		panic(err)
	}

	err = curproc.WriteCurrencies(config.OutputFile, currencies)
	if err != nil {
		panic(err)
	}
}
