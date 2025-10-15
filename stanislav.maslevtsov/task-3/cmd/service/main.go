package main

import (
	"fmt"
	"os"

	"github.com/jambii1/task-3/internal/configparser"
	"github.com/jambii1/task-3/internal/currenciesparser"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "-config" {
		fmt.Println("invalid command parameters")
	}

	config, err := configparser.ParseConfig(os.Args[2])
	if err != nil {
		panic(err)
	}

	currencies, err := currenciesparser.ParseCurrencies(config.InputFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(currencies.Сurncs[0].NumCode)
	fmt.Println(currencies.Сurncs[0].CharCode)
	fmt.Println(currencies.Сurncs[0].Value)
	fmt.Println(currencies.Сurncs[1].NumCode)
	fmt.Println(currencies.Сurncs[1].CharCode)
	fmt.Println(currencies.Сurncs[1].Value)
}
