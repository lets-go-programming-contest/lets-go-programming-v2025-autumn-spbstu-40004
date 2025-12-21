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

	fmt.Println(conf.Environment, conf.LogLevel)
}
