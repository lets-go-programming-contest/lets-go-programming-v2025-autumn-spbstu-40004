package main

import (
	"errors"
	"fmt"
)

const (
	minimalTemp = 15
	maximalTemp = 30
)

type TemperatureController struct {
	minTemp int
	maxTemp int
}

func NewTemperatureController() *TemperatureController {
	return &TemperatureController{
		minTemp: minimalTemp,
		maxTemp: maximalTemp,
	}
}

func (tc *TemperatureController) UpdateTemperature(operation string, personTemp int) error {
	switch operation {
	case ">=":
		if personTemp > tc.minTemp {
			tc.minTemp = personTemp
		}
	case "<=":
		if personTemp < tc.maxTemp {
			tc.maxTemp = personTemp
		}
	default:
		return errors.New("invalid operation")
	}

	return nil
}

func (tc *TemperatureController) GetCurrentTemperature() (int, error) {
	if tc.minTemp > tc.maxTemp {
		return -1, errors.New("temperature range is invalid")
	}
	return tc.minTemp, nil
}

func main() {
	var departmentCount int

	_, err := fmt.Scanln(&departmentCount)
	if err != nil {
		fmt.Printf("Error reading department count: %v\n", err)
		return
	}

	for range departmentCount {
		var peopleCount int

		_, err := fmt.Scanln(&peopleCount)
		if err != nil {
			fmt.Printf("Error reading people count: %v\n", err)
			return
		}

		tc := NewTemperatureController()

		for range peopleCount {
			var operation string
			var personTemp int

			_, err := fmt.Scan(&operation, &personTemp)
			if err != nil {
				fmt.Printf("Error reading operation or temperature: %v\n", err)
				return
			}

			err = tc.UpdateTemperature(operation, personTemp)
			if err != nil {
				fmt.Printf("Error processing department: %v\n", err)
				return
			}

			currentTemp, err := tc.GetCurrentTemperature()
			if err != nil {
				fmt.Println("-1")
			} else {
				fmt.Println(currentTemp)
			}
		}

	}
}
