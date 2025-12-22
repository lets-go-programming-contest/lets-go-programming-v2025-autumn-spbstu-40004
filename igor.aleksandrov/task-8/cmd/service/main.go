package main

import (
	"fmt"
	"os"

	"github.com/MrMels625/task-8/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration failure: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s %s\n", cfg.Environment, cfg.LogLevel)
}
