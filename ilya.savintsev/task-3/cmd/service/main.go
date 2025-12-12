package main

import (
	"flag"
	"os"

	filesaver "github.com/faxryzen/task-3/internal/file_saver"
	valsys "github.com/faxryzen/task-3/internal/valute_system"
	"gopkg.in/yaml.v2"
)

type DirHandle struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	var fileDir string

	flag.StringVar(&fileDir, "config", "yaml", "Specifies the path to the config")
	flag.Parse()

	content, err := os.ReadFile(fileDir)
	if err != nil {
		panic("no such file or directory")
	}

	var config DirHandle

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic("did not find expected key")
	}

	curs, err := valsys.ParseXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	jsonData, err := valsys.CreateJSON(curs)
	if err != nil {
		panic(err)
	}

	err = filesaver.SaveToFile(jsonData, config.OutputFile)
	if err != nil {
		panic(err)
	}
}
