package main

import (
	"fmt"
	"github.com/15446-rus75/task-8/config"
)

func main() {
	cfg := config.Load()
	fmt.Printf("%s %s\n", cfg.Environment, cfg.LogLevel)
}
