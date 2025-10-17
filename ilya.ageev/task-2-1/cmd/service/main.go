package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	minAllowedTemp = 15
	maxAllowedTemp = 30
)

type TemperatureController struct {
	minT int
	maxT int
}

func newTemperatureController() *TemperatureController {
	return &TemperatureController{
		minT: minAllowedTemp,
		maxT: maxAllowedTemp,
	}
}

func (tc *TemperatureController) changeMaxBound(currentTemp int) {
	if tc.minT == -1 {
		return
	}

	if currentTemp < tc.minT {
		tc.minT = -1
		return
	}

	if currentTemp <= tc.maxT {
		tc.maxT = currentTemp
	}
}

func (tc *TemperatureController) changeMinBound(currentTemp int) {
	if tc.minT == -1 {
		return
	}

	if currentTemp > tc.maxT {
		tc.minT = -1
		return
	}

	if currentTemp >= tc.minT {
		tc.minT = currentTemp
	}
}

func (tc *TemperatureController) findOptimalTemp(currentTemp int, str string) {
	switch str {
	case ">=":
		tc.changeMinBound(currentTemp)
	case "<=":
		tc.changeMaxBound(currentTemp)
	default:
		fmt.Println("Wrong input")

		return
	}
}

func (tc *TemperatureController) getTemperature() string {
	if tc.minT == -1 || tc.minT > tc.maxT {

		return "-1"
	}

	return strconv.Itoa(tc.minT)
}

func processDepartment(numWork int) []string {
	controller := newTemperatureController()
	departmentResults := make([]string, 0, numWork)

	for range numWork {

		var str string

		var currentTemp int

		_, err := fmt.Scan(&str, &currentTemp)
		if err != nil || currentTemp > maxAllowedTemp || currentTemp < minAllowedTemp {
			fmt.Println("Invalid temperature")

			return nil
		}

		controller.findOptimalTemp(currentTemp, str)
		optimalTemp := controller.getTemperature()
		departmentResults = append(departmentResults, optimalTemp)
	}

	return departmentResults
}

func main() {
	var numDepartments int

	_, err := fmt.Scan(&numDepartments)
	if err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	results := make([]string, 0)

	for range numDepartments {
		var numWork int

		_, err := fmt.Scan(&numWork)
		if err != nil {
			fmt.Println("invalid number of workers")

			return
		}

		departmentResults := processDepartment(numWork)
		if departmentResults == nil {
			return
		}

		results = append(results, departmentResults...)
	}

	fmt.Println(strings.Join(results, " "))
}
