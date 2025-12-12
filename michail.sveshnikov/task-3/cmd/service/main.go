package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/svesh3000/task-3/internal/centralbank"
	"github.com/svesh3000/task-3/internal/config"
)

func main() {
	configPath := flag.String("config", "./configs/config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("ERROR: Failed to load config: %v", err))
	}

	dir := filepath.Dir(cfg.OutputFile)

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("ERROR: Failed to create output directory: %v", err))
	}

	currencies, err := centralbank.ParseXMLFile(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("ERROR: Failed to parse XML file: %v", err))
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	err = centralbank.SaveCurrenciesToJSON(currencies, cfg.OutputFile)
	if err != nil {
		panic(fmt.Sprintf("ERROR: Failed to save JSON file: %v", err))
	}
}
