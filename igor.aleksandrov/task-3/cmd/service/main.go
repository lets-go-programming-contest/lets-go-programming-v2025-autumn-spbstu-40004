package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
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

	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic(fmt.Errorf("failed to read config file %s: %w", *configPath, err))
	}

	var conf Config

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal config file: %w", err))
	}

	fmt.Printf("Config loaded:\nInput: %s\nOutput: %s\n", conf.InputFile, conf.OutputFile)
}
