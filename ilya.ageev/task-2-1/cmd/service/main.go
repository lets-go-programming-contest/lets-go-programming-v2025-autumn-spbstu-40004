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

	_, err := fmt.Scan(&numOfDepart)
	if err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	results := make([]string, 0)

	for range numOfDepart {
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

		for range numOfWork {
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
				departmentResults = append(departmentResults, strconv.Itoa(minT))
			} else {
				departmentResults = append(departmentResults, "-1")
			}
		}

		results = append(results, departmentResults...)
	}

	fmt.Println(strings.Join(results, " "))
}
