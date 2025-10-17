package main

import (
	"errors"
	"fmt"
)

var errOperation = errors.New("Invalid operation")

func (minTemp int, maxTemp int, targetTemp int, operation string) (int, int, error) {
	if minTemp == -1 && maxTemp == -1 {
		return minTemp, maxTemp, nil
	}

	switch operation {
	case ">=":
		if targetTemp > maxTemp {
			minTemp = -1
			maxTemp = -1
		} else if minTemp <= targetTemp && targetTemp <= maxTemp {
			minTemp = targetTemp
		}
	case "<=":
		if targetTemp < minTemp {
			minTemp = -1
			maxTemp = -1
		} else if minTemp <= targetTemp && targetTemp <= maxTemp {
			maxTemp = targetTemp
		}
	default:
		return minTemp, maxTemp, errOperation
	}

	return minTemp, maxTemp, nil
}

func handleDepartmentRequests(employeeCount int, lowTemp int, hightTemp int) {
	var (
		condition  string
		targetTemp int
	)

	minTemp := lowTemp
	maxTemp := hightTemp

	for i := 0; i < employeeCount; i++ {
		_, err := fmt.Scanln(&condition, &targetTemp)
		if err != nil || targetTemp < 15 || targetTemp > 30 {
			fmt.Println("Invalid employee input")
			return
		}

		minTemp, maxTemp, err = updateTemperatureRange(minTemp, maxTemp, targetTemp, condition)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(minTemp)
	}
}

func main() {
	const (
		minTemperature = 15
		maxTemperature = 30
	)

	var departmentCount, employeeCount int

	_, err := fmt.Scanln(&departmentCount)
	if err != nil || departmentCount < 1 || departmentCount > 1000 {
		fmt.Println("Invalid department count")
		return
	}

	for i := 0; i < departmentCount; i++ {
		_, err = fmt.Scanln(&employeeCount)
		if err != nil || employeeCount < 1 || employeeCount > 1000 {
			fmt.Println("Invalid employee count")
			return
		}

		handleDepartmentRequests(employeeCount, minTemperature, maxTemperature)
	}
}
