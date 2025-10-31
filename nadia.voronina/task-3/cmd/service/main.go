package main

import (
	"github.com/alexflint/go-arg"
)

func main() {
	args := Args{}
	if err := arg.Parse(&args); err != nil {
		panic(err)
	}

	config, err := LoadConfig(args.Config)
	if err != nil {
		panic(err)
	}

	valCurs, err := ParseValuteXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	valJsons, err := ConvertValutesToJson(valCurs.Valutes)
	if err != nil {
		panic(err)
	}
	SortDescendingByValue(valJsons)

	if err := SaveToJson(valJsons, config.OutputFile); err != nil {
		panic(err)
	}
}
