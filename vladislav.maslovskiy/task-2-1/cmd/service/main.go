package main

import (
	"errors"
	"fmt"
)
var errOperator = errors.New("invalid operation")

func adjustTempurature(lowTempurature int, highTempurature int, askingTempurature int, operation string) (int, int, error) {
	if lowTempurature == -1 && highTempurature == -1 {
		return lowTempurature, highTempurature, nil
	}

	switch operation {
	case ">=":
		if askingTempurature > highTempurature {
			lowTempurature = -1
			highTempurature = -1
		} else if lowTempurature <= askingTempurature && askingTempurature <= highTempurature {
			lowTempurature = askingTempurature
		}
	case "<=":
		if askingTempurature < lowTempurature {
			lowTempurature = -1
			highTempurature = -1
		} else if lowTempurature <= askingTempurature && askingTempurature <= highTempurature {
			highTempurature = askingTempurature
		}
	default:
		return lowTempurature, highTempurature, errOperator
	}

	return lowTempurature, highTempurature, nil
}
func processDepart(employeeAmount int, minTempurature int, maxTempurature int) {
	var (
		operation  string
		askingTempurature int
	)

	lowTempurature := minTempurature
	highTempurature := maxTempurature
	for range employeeAmount {
		_, err := fmt.Scanln(&operation, &askingTempurature)
		if err != nil || askingTempurature < 15 || askingTempurature > 30 {
			fmt.Println("Wrong employee input")

			return
		}
		lowTempurature, highTempurature, err = adjustTempurature(lowTempurature, highTempurature, askingTempurature, operation)
		if err != nil {
			fmt.Println(err)

			return
		}
		fmt.Println(lowTempurature)
	}
}

func main() {
	const (
		minTempurature = 15
		highTempurature = 30
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
		processDepart(employeeAmount, minTempurature,highTempurature)
	}
}
