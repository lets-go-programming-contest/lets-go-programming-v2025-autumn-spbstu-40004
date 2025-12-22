package main

import (
	"fmt"

	"github.com/Ksenia-rgb/task-8/config"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s", conf.Environment, conf.LogLevel)
}
