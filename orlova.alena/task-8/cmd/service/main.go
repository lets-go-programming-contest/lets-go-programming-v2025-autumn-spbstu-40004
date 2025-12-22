package main

import (
	"fmt"

	"github.com/widgeiw/task-8/config"
)

func main() {
	cfg := config.GetConfig()
	fmt.Print(cfg.Environment, " ", cfg.LogLevel)
}
