package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"sort"

	"github.com/MrMels625/task-3/internal/config-utils"
	"github.com/MrMels625/task-3/internal/io-utils"
	"github.com/MrMels625/task-3/internal/xml-utils"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the YAML config file")

	flag.Parse()

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
