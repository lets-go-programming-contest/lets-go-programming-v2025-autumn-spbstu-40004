package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v2"
)

type configFile struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}
type valute struct {
	NumCode  int    `xml:"NumCode" json:"num_code"`
	CharCode string `xml:"CharCode" json:"char_code"`
	Value    string `xml:"Value" json:"-"`

	FloatValue float64 `xml:"-" json:"value"`
}

func (val *valute) convertFloatValue() error {
	val.Value = strings.ReplaceAll(val.Value, ",", ".")

	floatValue, err := strconv.ParseFloat(val.Value, 64)
	if err == nil {
		val.FloatValue = floatValue
	}
	return err
}

type inputFile struct {
	ValCurs []valute `xml:"Valute" json:"valute"`
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

	for index := range input.ValCurs {
		err := input.ValCurs[index].convertFloatValue()
		if err != nil {
			panic("Incorrect value format in input file")
		}
	}

	sort.Sort(input)

	splitPathOutput := strings.Split(config.Output, "/")
	if len(splitPathOutput) > 1 {
		dirName, _ := strings.CutSuffix(config.Output, splitPathOutput[len(splitPathOutput)-1])

		err := os.Mkdir(dirName, 0666)
		if err != nil && !os.IsExist(err) {
			panic("Incorrect permittion for output file")
		}
	}

	outputFileByte, err := json.MarshalIndent(input.ValCurs, "", "  ")
	if err != nil {
		panic("Error in output coding")
	}

	err = os.WriteFile(config.Output, outputFileByte, 0666)
	if err != nil {
		panic("Incorrect coding")
	}

}
