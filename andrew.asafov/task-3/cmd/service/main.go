package main

import (
	"flag"

	"github.com/shycoshy/task-3/internal/config"
	"github.com/shycoshy/task-3/internal/parser"
	"github.com/shycoshy/task-3/internal/sorter"
	"github.com/shycoshy/task-3/internal/writer"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("config file path is required")
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic("load config failed: " + err.Error())
	}

	data, err := parser.Parse(cfg.InputFile)
	if err != nil {
		panic("parse xml failed: " + err.Error())
	}

	sorted := sorter.Sort(data.Valutes)

	err = writer.Save(sorted, cfg.OutputFile)
	if err != nil {
		panic("save json failed: " + err.Error())
	}
}
