package main

import (
	"encoding/xml"
	"flag"
	"os"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v2"
)

type configFile struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}
type valute struct {
	NumCode   string
	CharCOde  string
	Nominal   int
	Name      string
	Value     string
	VunitRate string
}
type inputFile struct {
	ValCurs []valute `yaml:"Valute"`
}

func main() {
	flagConfig := flag.String("config", "congig.yaml", "direction config file for task3")
	flag.Parse()

	configFileByte, err := os.ReadFile(*flagConfig)
	if err != nil {
		panic("Incorrect config file")
	}

	config := configFile{}
	err = yaml.Unmarshal(configFileByte, &config)
	if err != nil || config.Input == "" || config.Output == "" {
		panic("Incorrect format in config file")
	}

	inputReader, err := os.Open(config.Input)
	if err != nil {
		panic("Incorrect input file")
	}

	decoder := xml.NewDecoder(inputReader)
	decoder.CharsetReader = charset.NewReaderLabel

	input := inputFile{}
	err = decoder.Decode(&input)
	if err != nil {
		panic("Incorrect formal in input file")
	}

}
