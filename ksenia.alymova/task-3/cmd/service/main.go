package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v2"
)

type configFile struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}
type valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func (val *valute) convertFloatValue() {
	(*val).Value = strings.ReplaceAll(val.Value, ",", ".")
}

type inputFile struct {
	ValCurs []valute `xml:"Valute"`
}

func (input inputFile) Len() int {
	return len(input.ValCurs)
}

func (input inputFile) Swap(index1, index2 int) {
	input.ValCurs[index1], input.ValCurs[index2] = input.ValCurs[index2], input.ValCurs[index1]
}

func (input inputFile) Less(index1, index2 int) bool {
	return input.ValCurs[index1].Value < input.ValCurs[index2].Value
}

func main() {
	flagConfig := flag.String("config", "congig.yaml", "direction config file for task3")
	flag.Parse()

	configFileByte, err := os.ReadFile(*flagConfig)
	if err != nil {
		panic("Incorrect config file")
	}

	var config configFile

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

	var input inputFile

	err = decoder.Decode(&input)
	if err != nil {
		panic("Incorrect formal in input file")
	}

	for index := 0; index < len(input.ValCurs); index++ {
		input.ValCurs[index].convertFloatValue()
	}

	sort.Sort(input)
	fmt.Println(input)
}
