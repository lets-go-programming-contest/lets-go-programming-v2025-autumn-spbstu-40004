package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"sort"

	configutils "github.com/MrMels625/task-3/internal/config-utils"
	ioutils "github.com/MrMels625/task-3/internal/io-utils"
	xmlutils "github.com/MrMels625/task-3/internal/xml-utils"
)

func main() {
	configPath := flag.String("config", "", "Path to the YAML config file")

	flag.Parse()

	if *configPath == "" {
		fmt.Println("Error: --config flag is required.")

		return
	}

	cfg, err := configutils.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := xmlutils.ParseXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	jsonMarshalled, err := json.MarshalIndent(currencies, "", "\t")
	if err != nil {
		panic(fmt.Errorf("json encoding error: %w", err))
	}

	err = ioutils.WriteBytesToFile(cfg.OutputFile, jsonMarshalled)
	if err != nil {
		panic(err)
	}
}
