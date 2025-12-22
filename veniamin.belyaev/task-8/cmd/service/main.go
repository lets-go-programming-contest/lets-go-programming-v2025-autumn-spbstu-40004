package main

import (
	"fmt"

	Config "github.com/belyaevEDU/task-8/internal/config"
)

func main() {
	config, err := Config.Parse()
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Print(config.Environment + " " + config.LogLevel)
}
