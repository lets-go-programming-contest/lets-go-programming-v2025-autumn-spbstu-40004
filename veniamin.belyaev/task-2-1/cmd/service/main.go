package main

import (
	"errors"
	"fmt"
)

var errOperation = errors.New("invalid operation")

func adjustTemperature(lowTemp int, highTemp int, askingTemp int, operation string) (int, int, error) {
	if lowTemp == -1 && highTemp == -1 {
		return lowTemp, highTemp, nil
	}

	switch operation {
	case ">=":
		if askingTemp > highTemp {
			lowTemp = -1
			highTemp = -1
		} else if lowTemp <= askingTemp && askingTemp <= highTemp {
			lowTemp = askingTemp
		}
	case "<=":
		if askingTemp < lowTemp {
			lowTemp = -1
			highTemp = -1
		} else if lowTemp <= askingTemp && askingTemp <= highTemp {
			highTemp = askingTemp
		}
	default:
		return lowTemp, highTemp, errOperation
	}

	return lowTemp, highTemp, nil
}

func processDepartment(employeeAmount int, minTemp int, maxTemp int) {
	var (
		operation  string
		askingTemp int
	)

	lowTemp := minTemp
	highTemp := maxTemp

	for range employeeAmount {
		_, err := fmt.Scanln(&operation, &askingTemp)
		if err != nil || askingTemp < 15 || askingTemp > 30 {
			fmt.Println("Invalid employee input")

			return
		}

		lowTemp, highTemp, err = adjustTemperature(lowTemp, highTemp, askingTemp, operation)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println(lowTemp)
	}
}

func main() {
	const (
		minTemp = 15
		maxTemp = 30
	)

	var departmentAmount, employeeAmount int

	_, err := fmt.Scanln(&departmentAmount)
	if err != nil || departmentAmount < 1 || departmentAmount > 1000 {
		fmt.Println("Invalid department amount")

		return
	}

	for range departmentAmount {
		_, err = fmt.Scanln(&employeeAmount)
		if err != nil || employeeAmount < 1 || employeeAmount > 1000 {
			fmt.Println("Invalid employee amount")

			return
		}

		processDepartment(employeeAmount, minTemp, maxTemp)
	}
}
