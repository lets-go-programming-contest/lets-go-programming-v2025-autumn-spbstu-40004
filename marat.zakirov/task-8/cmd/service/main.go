package main

import (
	"fmt"
	"log"

	"github.com/ZakirovMS/task-8/config"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(cfg.Environment, " ", cfg.LogLevel)
}
