package main

import (
	"flag"
	"fmt"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "", "Path to the YAML config file")

	flag.Parse()

	if *configPath == "" {
		fmt.Println("--config flag is required")

		return
	}

	fmt.Printf("Config file path: %s\n", *configPath)
}
