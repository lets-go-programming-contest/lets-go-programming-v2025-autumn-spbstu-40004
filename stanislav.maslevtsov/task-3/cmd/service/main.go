package main

import (
	"fmt"
	"os"

	"github.com/jambii1/task-3/internal/configparser"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "-config" {
		panic("configuration file not set")
	}

	config, err := configparser.ParseConfig(os.Args[2])
	if err != nil {
		panic(err)
	}

	fmt.Println(config.InputFile)
	fmt.Println(config.OutputFile)
}
