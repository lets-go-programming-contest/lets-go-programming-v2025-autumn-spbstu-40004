package main

import (
	"errors"
	"fmt"
)
var errOperator = errors.New("invalid operation")

func adjustTempurature(lowesttemp int, highesttemp int, askingTemp int, operation string) (int, int, error) {
	if lowesttemp == -1 && highesttemp == -1 {
		return lowesttemp, highesttemp, nil
	}

	switch operation {
	case ">=":
		if askingTemp > highesttemp {
			lowesttemp = -1
			highesttemp = -1
		} else if lowesttemp <= askingTemp && askingTemp <= highesttemp {
			lowesttemp = askingTemp
		}
	case "<=":
		if askingTemp < lowesttemp {
			lowesttemp = -1
			highesttemp = -1
		} else if lowesttemp <= askingTemp && askingTemp <= highesttemp {
			highesttemp = askingTemp
		}
	default:
		return lowesttemp, highesttemp, errOperator
	}

	return lowesttemp, highesttemp, nil
}

func processDepart(employeeAmount int, minTempurature int, maxTempurature int) {
	var (
		operation  string
		askingTemp int
	)

	lowesttemp := minTempurature
	highesttemp := maxTempurature

	for range employeeAmount {
		_, err := fmt.Scanln(&operation, &askingTemp)
		if err != nil || askingTemp < 15 || askingTemp > 30 {
			fmt.Println("Wrong employee input")

			return
		}

		lowesttemp, highesttemp, err = adjustTempurature(lowesttemp, highesttemp, askingTemp, operation)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println(lowesttemp)
	}
}

func main() {
	const (
		minTempurature = 15
		highesttemp = 30
	)
	var departmentAmount, employeeAmount int
	_, err := fmt.Scanln(&departmentAmount)
	if err != nil || departmentAmount < 1 || departmentAmount > 1000 {
		fmt.Println("Wrong  department amount")

		return
	}

	for range departmentAmount {
		_, err = fmt.Scanln(&employeeAmount)
		if err != nil || employeeAmount < 1 || employeeAmount > 1000 {
			fmt.Println("Wrong employee amount")

			return
		}
		processDepart(employeeAmount, minTempurature,highesttemp)
	}
}
