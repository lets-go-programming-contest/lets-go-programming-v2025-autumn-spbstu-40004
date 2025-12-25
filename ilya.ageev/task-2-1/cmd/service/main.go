package main

import (
	"fmt"
	"strconv"
	"strings"
)

func processDepartment(numOfWork int) []string {
	const (
		minAllowedTemp = 15
		maxAllowedTemp = 30
	)

	minT := 15
	maxT := 30
	departmentResults := make([]string, 0, numOfWork)

	for range numOfWork {
		var str string

		var Temp int

		_, err := fmt.Scan(&str, &Temp)
		if err != nil || Temp > maxAllowedTemp || Temp < minAllowedTemp {
			fmt.Println("Invalid temperature")

			return nil
		}

		if str == ">=" {
			if Temp > minT {
				minT = Temp
			}
		} else if str == "<=" {
			if Temp < maxT {
				maxT = Temp
			}
		}

		if minT <= maxT {
			departmentResults = append(departmentResults, strconv.Itoa(minT))
		} else {
			departmentResults = append(departmentResults, "-1")
		}
	}

	return departmentResults
}

func main() {
	var numOfDepart int

	_, err := fmt.Scan(&numOfDepart)
	if err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	results := make([]string, 0)

	for range numOfDepart {
		var numOfWork int

		_, err := fmt.Scan(&numOfWork)
		if err != nil {
			fmt.Println("invalid number of workers")

			return
		}

		departmentResults := processDepartment(numOfWork)
		if departmentResults == nil {
			return
		}

		results = append(results, departmentResults...)
	}

	fmt.Println(strings.Join(results, "\n"))
}
