package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		n, k, T int
		str     string
	)
	results := make([]string, 0)

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid number of departments")
		return
	}

	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&k)
		if err != nil {
			fmt.Println("invalid number of workers")
			return
		}

		minT := 15
		maxT := 30
		departmentResults := make([]string, 0)

		for j := 0; j < k; j++ {
			_, err := fmt.Scan(&str, &T)
			if err != nil {
				fmt.Println("Invalid temperature")
				return
			}

			if str == ">=" {
				if T > minT {
					minT = T
				}
			} else if str == "<=" {
				if T < maxT {
					maxT = T
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
