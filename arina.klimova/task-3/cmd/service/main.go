package main

import (
	"fmt"

	"github.com/arinaklimova/task-3/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Printf("Configuration loaded successfully:")
	fmt.Printf("Input file: %s\n", cfg.InputFile)
	fmt.Printf("Output file: %s\n", cfg.OutputFile)
}
