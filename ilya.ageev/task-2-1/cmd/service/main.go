package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		n, k, minT, maxT, T int
		str                 string
	)
	results := make([]string, 0)
	minT = 15
	maxT = 30
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid number of departments")
	}
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&k)
		if err != nil {
			fmt.Println("invalid number of workers")
		}
		for j := 0; j < k; j++ {
			_, err := fmt.Scan(&str, &T)
			if err != nil {
				fmt.Println("Invalid temperature")
			}
			if str == ">=" {
				if T > minT {
					minT = T
					fmt.Println("minT: ", minT)
				}
			} else if str == "<=" {
				if T < maxT {
					maxT = T
					fmt.Println("maxT: ", maxT)
				}
			}

			if minT <= maxT {
				results = append(results, strconv.Itoa(minT))
			} else {
				results = append(results, "-1")
			}
		}
		minT = 15
		maxT = 30
	}
	fmt.Println(strings.Join(results, " "))
}
