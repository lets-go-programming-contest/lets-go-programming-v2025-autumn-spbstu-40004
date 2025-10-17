package main

import (
	"fmt"
)

const (
	minimalTemp = 15
	maximalTemp = 30
)

var (
	ErrReadingDepartment = fmt.Errorf("error reading department count")
	ErrReadingPeople     = fmt.Errorf("error reading people count")
	ErrReadingTemp       = fmt.Errorf("error reading operation or temperature")
	ErrProcessingDept    = fmt.Errorf("error processing department")
	ErrInvalidOperation  = fmt.Errorf("invalid operation")
	ErrInvalidTempRange  = fmt.Errorf("temperature range is invalid")
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
		return ErrInvalidOperation
	}

	return nil
}

func (tc *TemperatureController) GetCurrentTemperature() (int, error) {
	if tc.minTemp > tc.maxTemp {
		return -1, ErrInvalidTempRange
	}

	return tc.minTemp, nil
}

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

		tempControl := NewTemperatureController()

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
