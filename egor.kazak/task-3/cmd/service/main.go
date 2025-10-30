package main

import (
	"encoding/json"
	"log"

	"github.com/CuatHimBong/task-3/internal/config"
	"github.com/CuatHimBong/task-3/internal/currency"
	"github.com/CuatHimBong/task-3/internal/storage"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	xmlData, err := storage.ReadFile(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.ParseXML(xmlData)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		panic(err)
	}

	if err := storage.WriteJSON(cfg.OutputFile, jsonData); err != nil {
		panic(err)
	}

	log.Println("Success: data written to", cfg.OutputFile)
}
