package main

import (
	"errors"
	"fmt"
)

const (
	minAbsoluteTemperature = 15
	maxAbsoluteTemperature = 30
	invalidTemperature     = -1
)

var (
	ErrInvalidComparisonSign       = errors.New("unsupported comparison sign for temperature")
	ErrUnsupportedTemperature      = errors.New("unsupported temperature value")
	ErrInvalidTemperatureValue     = errors.New("invalid temperature value format")
	ErrInvalidDepartmentCount      = errors.New("invalid departments count")
	ErrInvalidEmployeesCount       = errors.New("invalid employees count")
	ErrInvalidComparisonSignFormat = errors.New("invalid comparison sign format for temperature")
)

type OptimalTemperature struct {
	min int
	max int
}

func NewOptimalTemperature() OptimalTemperature {
	return OptimalTemperature{
		min: minAbsoluteTemperature,
		max: maxAbsoluteTemperature,
	}
}

func (ot *OptimalTemperature) Update(comparisonSign string, temperature int) error {
	if temperature < minAbsoluteTemperature || temperature > maxAbsoluteTemperature {
		return ErrUnsupportedTemperature
	}

	switch comparisonSign {
	case ">=":
		if temperature > ot.min {
			ot.min = temperature
		}
	case "<=":
		if temperature < ot.max {
			ot.max = temperature
		}
	default:
		return ErrInvalidComparisonSign
	}

	return nil
}

func (ot *OptimalTemperature) GetOptimal() int {
	if ot.min > ot.max {
		return invalidTemperature
	}

	return ot.min
}

func processEmployeesData(employeesCount int) {
	optimalTemp := NewOptimalTemperature()

	for range employeesCount {
		var comparisonSign string

		var temperature int

		_, err := fmt.Scan(&comparisonSign)
		if err != nil {
			fmt.Println(ErrInvalidComparisonSignFormat.Error())

			continue
		}

		_, err = fmt.Scan(&temperature)
		if err != nil {
			fmt.Println(ErrInvalidTemperatureValue.Error())

			continue
		}

		err = optimalTemp.Update(comparisonSign, temperature)
		if err != nil {
			fmt.Println(err.Error())

			continue
		}

		fmt.Println(optimalTemp.GetOptimal())
	}
}

func main() {
	var departmentsCount int

	_, err := fmt.Scanln(&departmentsCount)
	if err != nil || departmentsCount < 1 || departmentsCount > 1000 {
		fmt.Println(ErrInvalidDepartmentCount.Error())

		return
	}

	for range departmentsCount {
		var employeesCount int

		_, err = fmt.Scanln(&employeesCount)
		if err != nil || employeesCount < 1 || employeesCount > 1000 {
			fmt.Println(ErrInvalidEmployeesCount.Error())

			continue
		}

		processEmployeesData(employeesCount)
	}
}
