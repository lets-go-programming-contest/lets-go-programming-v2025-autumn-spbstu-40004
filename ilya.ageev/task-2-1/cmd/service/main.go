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
	minT  int
	maxT  int
	valid bool
}

func newTemperatureController() *TemperatureController {
	return &TemperatureController{
		minT:  minAllowedTemp,
		maxT:  maxAllowedTemp,
		valid: true,
	}
}

func (tc *TemperatureController) applyConstraint(op string, temp int) {
	if !tc.valid {
		return
	}

	switch op {
	case ">=":
		if temp > tc.maxT {
			tc.valid = false
		} else if temp > tc.minT {
			tc.minT = temp
		}
	case "<=":
		if temp < tc.minT {
			tc.valid = false
		} else if temp < tc.maxT {
			tc.maxT = temp
		}
	}

	if tc.minT > tc.maxT {
		tc.valid = false
	}
}

func (tc *TemperatureController) getTemperature() string {
	if !tc.valid {
		return "-1"
	}
	return strconv.Itoa(tc.minT)
}

func processDepartment(numWork int) []string {
	controller := newTemperatureController()
	departmentResults := make([]string, 0, numWork)

	for i := 0; i < numWork; i++ {
		var op string
		var temp int

		_, err := fmt.Scan(&op, &temp)
		if err != nil {
			controller.valid = false
		} else if temp < minAllowedTemp || temp > maxAllowedTemp {
			controller.valid = false
		} else {
			controller.applyConstraint(op, temp)
		}

		departmentResults = append(departmentResults, controller.getTemperature())
	}

	return departmentResults
}

func main() {
	var numDepartments int

	_, err := fmt.Scan(&numDepartments)
	if err != nil {
		return
	}

	allResults := make([]string, 0)

	for i := 0; i < numDepartments; i++ {
		var numWork int

		_, err := fmt.Scan(&numWork)
		if err != nil {
			return
		}

		departmentResults := processDepartment(numWork)
		allResults = append(allResults, departmentResults...)
	}

	fmt.Println(strings.Join(allResults, " "))
}
