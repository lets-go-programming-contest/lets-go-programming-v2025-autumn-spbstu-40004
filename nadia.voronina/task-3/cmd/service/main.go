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
	SortDescendingByValue(&valCurs)

	if err := SaveToJson(valCurs.Valutes, config.OutputFile); err != nil {
		panic(err)
	}
}
