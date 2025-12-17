package main

import (
	"fmt"

	"github.com/jambii1/task-8/pkg/config"
)

func main() {
	conf, err := config.Parse()
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Print(conf.Environment + " " + conf.LogLevel)
}
