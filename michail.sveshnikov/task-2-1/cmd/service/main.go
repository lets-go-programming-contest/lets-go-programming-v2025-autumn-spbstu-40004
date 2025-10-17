package main

import (
	"fmt"

	"github.com/svesh3000/task-2-1/internal/conditioner"
)

func main() {
	const (
		lowLimit int = 15
		upLimit  int = 30
	)

	var depNum int

	_, err := fmt.Scanln(&depNum)
	if err != nil || depNum < 1 || depNum > 1000 {
		fmt.Println("ERROR: invalid number of departments!")

		return
	}

	for range depNum {
		var employeeNum int

		_, err = fmt.Scanln(&employeeNum)
		if err != nil || employeeNum < 1 || employeeNum > 1000 {
			fmt.Println("ERROR: invalid number of employees!")

			continue
		}

		cond := conditioner.Conditioner{
			LowBound: lowLimit,
			UpBound:  upLimit,
		}

		for range employeeNum {
			var (
				temperature int
				compSign    string
			)

			_, err = fmt.Scanln(&compSign, &temperature)
			if err != nil {
				fmt.Println("ERROR: invalid request!")

				continue
			}

			optimizeTemp, err := cond.OptimizeTemperature(compSign, temperature)
			if err != nil {
				fmt.Println(err)

				continue
			}

			fmt.Println(optimizeTemp)
		}
	}
}
