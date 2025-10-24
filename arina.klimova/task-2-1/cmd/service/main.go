package main

import (
	"errors"
	"fmt"

	TempControl "github.com/arinaklimova/task-2-1/internal/temperature"
)

var (
	ErrReadingDepartment = errors.New("error reading department count")
	ErrReadingPeople     = errors.New("error reading people count")
	ErrReadingTemp       = errors.New("error reading operation or temperature")
	ErrProcessingDept    = errors.New("error processing department")
)

func main() {
	var departmentCount int

	_, err := fmt.Scanln(&departmentCount)
	if err != nil {
		fmt.Printf("%v: %v\n", ErrReadingDepartment, err)

		return
	}

	for range departmentCount {
		var peopleCount int

		_, err := fmt.Scanln(&peopleCount)
		if err != nil {
			fmt.Printf("%v: %v\n", ErrReadingPeople, err)

			return
		}

		tempControl := TempControl.NewTemperatureController()

		for range peopleCount {
			var (
				operation  string
				personTemp int
			)

			_, err := fmt.Scan(&operation, &personTemp)
			if err != nil {
				fmt.Printf("%v: %v\n", ErrReadingTemp, err)

				return
			}

			err = tempControl.UpdateTemperature(operation, personTemp)
			if err != nil {
				fmt.Printf("%v: %v\n", ErrProcessingDept, err)

				return
			}

			currentTemp, err := tempControl.GetCurrentTemperature()
			if err != nil {
				fmt.Println("-1")
			} else {
				fmt.Println(currentTemp)
			}
		}
	}
}
