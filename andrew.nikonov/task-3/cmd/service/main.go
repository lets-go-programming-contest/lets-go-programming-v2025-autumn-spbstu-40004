package main

import (
	"flag"

	"github.com/ysffmn/task-3/internal/config"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("Config flag is required")
	}

	cfg := config.Load(*configPath)
	_ = cfg
}
