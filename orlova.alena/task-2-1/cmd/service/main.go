package main

import (
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

func main() {
	var (
		currTemp                       int
		numOfDepartments, numOfWorkers int
		sign                           string
	)

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
			_, err = fmt.Scanln(&sign, &currTemp)
			if err != nil {
				fmt.Println("Wrong input")

				return
			}

			switch sign {
			case ">=":
				if currTemp > maxTempBound {
					fmt.Println("-1")

					return
				}

				if currTemp >= minTempBound {
					minTempBound = currTemp
				}
			case "<=":
				if currTemp < minTempBound {
					fmt.Println("-1")

					return
				}

				if currTemp <= maxTempBound {
					maxTempBound = currTemp
				}
			default:
				fmt.Println("Wrong input")

				return
			}

			fmt.Println(minTempBound)
		}
	}
}
