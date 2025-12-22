package main

import (
	"fmt"

	"github.com/MrMels625/task-8/config"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		return
	}

	fmt.Print(conf.Environment, " ", conf.LogLevel)
}
