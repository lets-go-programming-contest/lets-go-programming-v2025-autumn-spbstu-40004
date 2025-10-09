package main

import (
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

func changeMaxBound(minTempBound, maxTempBound *int, currTemp int) {
	if *minTempBound == -1 {
		return
	}

	if currTemp < *minTempBound {
		*minTempBound = -1

		return
	}

	if currTemp <= *maxTempBound {
		*maxTempBound = currTemp
	}
}

func changeMinBound(minTempBound, maxTempBound *int, currTemp int) {
	if *minTempBound == -1 {
		return
	}

	if currTemp > *maxTempBound {
		*minTempBound = -1

		return
	}

	if currTemp >= *minTempBound {
		*minTempBound = currTemp
	}
}

func findOptimalTemp(minTempBound, maxTempBound *int, currTemp int, sign string) {
	switch sign {
	case ">=":
		changeMinBound(minTempBound, maxTempBound, currTemp)
	case "<=":
		changeMaxBound(minTempBound, maxTempBound, currTemp)
	default:
		fmt.Println("Wrong input")

		return
	}
}

func main() {
	var (
		numOfDeparts, numOfWorkers, currTemp int
		sign                                 string
	)

	_, err := fmt.Scanln(&numOfDeparts)
	if err != nil {
		fmt.Println("Wrong input")

		return
	}

	for i := 1; i <= numOfDeparts; i++ {
		_, err = fmt.Scanln(&numOfWorkers)
		if err != nil {
			fmt.Println("Wrong input")

			return
		}

		minTempBound := minTemp
		maxTempBound := maxTemp

		for j := 1; j <= numOfWorkers; j++ {
			_, err = fmt.Scanln(&sign, &currTemp)
			if err != nil {
				fmt.Println("Wrong input")

				return
			}

			findOptimalTemp(&minTempBound, &maxTempBound, currTemp, sign)
			fmt.Println(minTempBound)
		}
	}
}
