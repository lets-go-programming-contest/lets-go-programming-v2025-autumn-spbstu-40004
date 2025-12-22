package main

import (
	"fmt"

	"github.com/ZakirovMS/task-8/config"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s\n", cfg.Environment, cfg.LogLevel)
}
