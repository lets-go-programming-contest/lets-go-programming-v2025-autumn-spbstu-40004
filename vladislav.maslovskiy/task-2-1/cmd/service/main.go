package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidOperation = errors.New("invalid operation")
	ErrWrongDepartment  = errors.New("wrong department amount")
	ErrWrongEmployee    = errors.New("wrong employee amount")
	ErrWrongInput       = errors.New("wrong employee input")
	ErrInvalidTemp      = errors.New("invalid temperature")
)

const (
	MinTemperature = 15
	MaxTemperature = 30
	MinDepartments = 1
	MaxDepartments = 1000
	MinEmployees   = 1
	MaxEmployees   = 1000
)

type TemperatureRange struct {
	lowest  int
	highest int
}

func NewTemperatureRange() *TemperatureRange {
	return &TemperatureRange{
		lowest:  MinTemperature,
		highest: MaxTemperature,
	}
}

func (tempRange *TemperatureRange) Adjust(askingTemp int, operation string) error {
	if tempRange.lowest == -1 && tempRange.highest == -1 {
		return nil
	}

	switch operation {
	case ">=":
		if askingTemp > tempRange.highest {
			tempRange.lowest = -1
			tempRange.highest = -1
		} else if tempRange.lowest <= askingTemp && askingTemp <= tempRange.highest {
			tempRange.lowest = askingTemp
		}
	case "<=":
		if askingTemp < tempRange.lowest {
			tempRange.lowest = -1
			tempRange.highest = -1
		} else if tempRange.lowest <= askingTemp && askingTemp <= tempRange.highest {
			tempRange.highest = askingTemp
		}
	default:
		return ErrInvalidOperation
	}

	return nil
}

func (tempRange *TemperatureRange) GetCurrent() int {
	if tempRange.lowest == -1 && tempRange.highest == -1 {
		return -1
	}

	return tempRange.lowest
}

func ProcessDepartment(employeeAmount int, requests []struct {
	operation  string
	askingTemp int
}) ([]int, error) {
	temperatureRange := NewTemperatureRange()
	results := make([]int, 0, employeeAmount)

	for i := range employeeAmount {
		req := requests[i]

		if req.askingTemp < MinTemperature || req.askingTemp > MaxTemperature {
			return nil, ErrInvalidTemp
		}

		err := temperatureRange.Adjust(req.askingTemp, req.operation)
		if err != nil {
			return nil, err
		}

		results = append(results, temperatureRange.GetCurrent())
	}

	return results, nil
}

func main() {
	var departmentAmount int

	_, err := fmt.Scanln(&departmentAmount)
	if err != nil || departmentAmount < MinDepartments || departmentAmount > MaxDepartments {
		fmt.Println(ErrWrongDepartment)

		return
	}

	for range departmentAmount {
		var employeeAmount int

		_, err := fmt.Scanln(&employeeAmount)
		if err != nil || employeeAmount < MinEmployees || employeeAmount > MaxEmployees {
			fmt.Println(ErrWrongEmployee)

			return
		}

		requests := make([]struct {
			operation  string
			askingTemp int
		}, employeeAmount)

		for j := range employeeAmount {
			_, err := fmt.Scanln(&requests[j].operation, &requests[j].askingTemp)
			if err != nil {
				fmt.Println(ErrWrongInput)

				return
			}
		}

		results, err := ProcessDepartment(employeeAmount, requests)
		if err != nil {
			fmt.Println(err)

			return
		}

		for _, result := range results {
			fmt.Println(result)
		}
	}
}

