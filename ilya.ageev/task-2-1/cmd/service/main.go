package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		numOfDepart, numOfWork, Temp int
		str                          string
	)
	results := make([]string, 0)

	_, err := fmt.Scan(&numOfDepart)
	if err != nil {
		fmt.Println("Invalid number of departments")
		return
	}

	for i := 0; i < numOfDepart; i++ {
		_, err := fmt.Scan(&numOfWork)
		if err != nil {
			fmt.Println("invalid number of workers")
			return
		}
		const (
			minAllowedTemp = 15
			maxAllowedTemp = 30
		)
		minT := 15
		maxT := 30
		departmentResults := make([]string, 0)

		for j := 0; j < numOfWork; j++ {
			_, err := fmt.Scan(&str, &Temp)
			if err != nil || Temp > maxAllowedTemp || Temp < minAllowedTemp {
				fmt.Println("Invalid temperature")
				return
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
				// Use the maximum possible temperature that satisfies all constraints
				departmentResults = append(departmentResults, strconv.Itoa(minT))
			} else {
				departmentResults = append(departmentResults, "-1")
			}
		}

		// Add all results for this department to the main results
		results = append(results, departmentResults...)
	}

	fmt.Println(strings.Join(results, " "))
}
