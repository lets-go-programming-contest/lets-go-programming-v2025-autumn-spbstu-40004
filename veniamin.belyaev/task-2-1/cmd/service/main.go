package main

import (
	"errors"
	"fmt"
)

var errOperation = errors.New("invalid operation")

// cyclomatic complexity workaround
func isValueInInterval(value int, min int, max int) bool {
	return min <= value && value <= max
}

func adjustTemperature(lowTemp int, highTemp int, askingTemp int, operation string) (int, int, error) {
	if lowTemp == -1 && highTemp == -1 {
		return lowTemp, highTemp, nil
	}

	switch operation {
	case ">=":
		if askingTemp > highTemp {
			lowTemp = -1
			highTemp = -1
		} else if isValueInInterval(lowTemp, askingTemp, highTemp) {
			lowTemp = askingTemp
		}
	case "<=":
		if askingTemp < lowTemp {
			lowTemp = -1
			highTemp = -1
		} else if isValueInInterval(lowTemp, askingTemp, highTemp) {
			highTemp = askingTemp
		}
	default:
		return lowTemp, highTemp, errOperation
	}

	return lowTemp, highTemp, nil
}

func main() {
	const (
		minTemp = 15
		maxTemp = 30
	)

	var (
		departmentAmount, employeeAmount, askingTemp int
		operation                                    string
	)

	_, err := fmt.Scanln(&departmentAmount)
	if err != nil || !isValueInInterval(departmentAmount, 1, 1000) {
		fmt.Println("Invalid department amount")

		return
	}

	for range departmentAmount {
		_, err = fmt.Scanln(&employeeAmount)
		if err != nil || !isValueInInterval(employeeAmount, 1, 1000) {
			fmt.Println("Invalid employee amount")

			return
		}

		lowTemp := minTemp
		highTemp := maxTemp

		for range employeeAmount {
			_, err = fmt.Scanln(&operation, &askingTemp)
			if err != nil || !isValueInInterval(askingTemp, 15, 30) {
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
}
