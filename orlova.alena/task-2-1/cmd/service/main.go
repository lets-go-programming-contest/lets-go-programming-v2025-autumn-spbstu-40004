package main

import (
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

func findOptimalTemp(minTempBound, maxTempBound *int) {
	var (
		currTemp int
		sign     string
	)
	_, err := fmt.Scanln(&sign, &currTemp)
	if err != nil {
		fmt.Println("Wrong input")

		return
	}

	switch sign {
	case ">=":
		if *minTempBound == -1 {
			return
		} else if currTemp > *maxTempBound {
			*minTempBound = -1
		} else if currTemp >= *minTempBound {
			*minTempBound = currTemp
		}
	case "<=":
		if *minTempBound == -1 {
			return
		} else if currTemp < *minTempBound {
			*minTempBound = -1
		} else if currTemp <= *maxTempBound {
			*maxTempBound = currTemp
		}
	default:
		fmt.Println("Wrong input")

		return
	}
}

func main() {
	var numOfDepartments, numOfWorkers int
	_, err := fmt.Scanln(&numOfDepartments)
	if err != nil {
		fmt.Println("Wrong input")

		return
	}

	for i := 1; i <= numOfDepartments; i++ {
		_, err = fmt.Scanln(&numOfWorkers)
		if err != nil {
			fmt.Println("Wrong input")

			return
		}

		minTempBound := minTemp
		maxTempBound := maxTemp
		for j := 1; j <= numOfWorkers; j++ {
			findOptimalTemp(&minTempBound, &maxTempBound)
			fmt.Println(minTempBound)
		}
	}
}
