package main

import (
	"flag"
	"sort"

	"github.com/ysffmn/task-3/internal/config"
	"github.com/ysffmn/task-3/internal/currency"
	"github.com/ysffmn/task-3/internal/parser"
	"github.com/ysffmn/task-3/internal/writer"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	var valCurs currency.ValCurs

	err = parser.Parse(cfg.InputFile, &valCurs)
	if err != nil {
		panic(err)
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	err = writer.Write(cfg.OutputFile, valCurs.Valutes)
	if err != nil {
		panic(err)
	}
}
