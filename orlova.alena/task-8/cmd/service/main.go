package main

import (
	"fmt"

	"github.com/widgeiw/task-8/config"
)

func main() {
	cfg := config.GetConfig()
	fmt.Printf("%s %s\n", cfg.Environment, cfg.LogLevel)
}
