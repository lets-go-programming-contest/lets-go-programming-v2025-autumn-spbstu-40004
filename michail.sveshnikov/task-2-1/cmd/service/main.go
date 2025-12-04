package main

import (
	"errors"
	"fmt"

	"github.com/svesh3000/task-2-1/internal/conditioner"
)

var (
	errInvalidDepNum      = errors.New("invalid number of departments")
	errInvalidEmployeeNum = errors.New("invalid number of employees")
	errInvalidRequest     = errors.New("invalid request")
)

func main() {
	const (
		lowTempLimit   = 15
		upTempLimit    = 30
		minDepNum      = 1
		maxDepNum      = 1000
		minEmployeeNum = 1
		maxEmployeeNum = 1000
	)

	var depNum int

	_, err := fmt.Scanln(&depNum)
	if err != nil || depNum < minDepNum || depNum > maxDepNum {
		fmt.Println(errInvalidDepNum)

		return
	}

	for range depNum {
		var employeeNum int

		_, err = fmt.Scanln(&employeeNum)
		if err != nil || employeeNum < minEmployeeNum || employeeNum > maxEmployeeNum {
			fmt.Println(errInvalidEmployeeNum)

			continue
		}

		cond := conditioner.Conditioner{
			LowBound: lowTempLimit,
			UpBound:  upTempLimit,
		}

		for range employeeNum {
			var (
				temperature int
				compSign    string
			)

			_, err = fmt.Scanln(&compSign, &temperature)
			if err != nil {
				fmt.Println(errInvalidRequest)

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
