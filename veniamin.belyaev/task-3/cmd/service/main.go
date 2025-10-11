package main

import (
	"flag"
	"strings"

	config_handler "github.com/belyaevEDU/task-3/internal"
)

func main() {
	var configFilePath string

	flag.StringVar(&configFilePath, "config", "none", "Configuration file path")
	flag.Parse()

	if strings.Compare(configFilePath, "") == 0 {
		panic("config file path via flag not specified")
	}

	config, err := config_handler.LoadConfig(configFilePath)
	if err != nil {
		panic(err.Error())
	}

}
