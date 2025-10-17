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

	if currentTemp < tc.maxT {
		tc.maxT = currentTemp
	}

	if tc.minT > tc.maxT {
		tc.minT = -1
		tc.maxT = -1
	}
}

func (tc *TemperatureController) changeMinBound(currentTemp int) {
	if tc.minT == -1 {
		return
	}

	if currentTemp > tc.minT {
		tc.minT = currentTemp
	}

	if tc.minT > tc.maxT {
		tc.minT = -1
		tc.maxT = -1
	}
}

func (tc *TemperatureController) findOptimalTemp(currentTemp int, str string) {
	switch str {
	case ">=":
		tc.changeMinBound(currentTemp)
	case "<=":
		tc.changeMaxBound(currentTemp)
	}
}

func (tc *TemperatureController) getTemperature() string {
	if tc.minT == -1 {
		return "-1"
	}
	return strconv.Itoa(tc.minT)
}

func processDepartment(numWork int) []string {
	controller := newTemperatureController()
	departmentResults := make([]string, 0, numWork)

	for i := 0; i < numWork; i++ {
		var str string
		var currentTemp int

		_, err := fmt.Scan(&str, &currentTemp)
		if err != nil {
			controller = newTemperatureController()
		}

		// Validate temperature range
		if currentTemp < minAllowedTemp || currentTemp > maxAllowedTemp {
			controller.minT = -1
			controller.maxT = -1
		} else {
			controller.findOptimalTemp(currentTemp, str)
		}

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

	for i := 0; i < numDepartments; i++ {
		var numWork int

		_, err := fmt.Scan(&numWork)
		if err != nil {
			fmt.Println("Invalid number of workers")
			return
		}

		departmentResults := processDepartment(numWork)
		results = append(results, departmentResults...)
	}

	fmt.Println(strings.Join(results, " "))
}
