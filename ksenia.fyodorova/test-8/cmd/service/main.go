package main

import (
	"fmt"
	"os"

	"github.com/lolnyok/task-8/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки конфигурации: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s %s\n", cfg.Environment, cfg.LogLevel)
}
