package main

import (
	"github.com/alexflint/go-arg"
	"spbstu.ru/nadia.voronina/task-3/pkg/args"
	"spbstu.ru/nadia.voronina/task-3/pkg/config"
	"spbstu.ru/nadia.voronina/task-3/pkg/sort"
	"spbstu.ru/nadia.voronina/task-3/pkg/valute"
)

func main() {
	args := args.Args{Config: ""}
	if err := arg.Parse(&args); err != nil {
		panic(err)
	}

	config, err := config.LoadConfig(args.Config)
	if err != nil {
		panic(err)
	}

	valCurs, err := valute.ParseValuteXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	sort.SortDescendingByValue(valCurs.Valutes)

	valJsons, err := valute.ConvertValutesToJSONBytes(valCurs.Valutes)
	if err != nil {
		panic(err)
	}

	if err := valute.SaveJSONBytes(valJsons, config.OutputFile); err != nil {
		panic(err)
	}
}
