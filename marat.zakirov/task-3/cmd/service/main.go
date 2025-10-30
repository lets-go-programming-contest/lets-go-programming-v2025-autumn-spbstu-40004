package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ZakirovMS/task-3/internal/pathHolder"
	"gopkg.in/yaml.v3"
)

func main() {
	var nFlag = flag.String("config", "", "Path to YAML config file")

	flag.Parse()
	configFile, err := os.ReadFile(*nFlag)
	if err != nil {
		panic("Some error in getting config file")
	}

	var ioPath pathHolder.PathHolder
	err = yaml.Unmarshal(configFile, &ioPath)
	if err != nil {
		panic("Some error in reading config file")
	}

	fmt.Println(ioPath)
}
