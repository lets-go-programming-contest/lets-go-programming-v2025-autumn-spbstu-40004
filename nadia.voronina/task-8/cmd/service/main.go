package main

import (
	"fmt"

	"spbstu.ru/nadia.voronina/task-8/config"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s", conf.Environment, conf.LogLevel)
}
