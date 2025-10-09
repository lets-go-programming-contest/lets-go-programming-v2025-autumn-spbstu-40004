package main

import "fmt"

const (
	minAllowedTemp = 15
	maxAllowedTemp = 30
)

func processDept(staffCount int) {
	minTemp := minAllowedTemp
	maxTemp := maxAllowedTemp

	for range staffCount {
		var (
			operation string
			temp      int
		)

		_, err := fmt.Scanln(&operation, &temp)
		if err != nil || temp > maxAllowedTemp || temp < minAllowedTemp {
			fmt.Println("Invalid temperature number")

			continue
		}

		switch operation {
		case ">=":
			if temp > minTemp {
				minTemp = temp
			}
		case "<=":
			if temp < maxTemp {
				maxTemp = temp
			}
		default:
			fmt.Println("Invalid operation")

			continue
		}

		if minTemp > maxTemp {
			fmt.Println("-1")
		} else {
			fmt.Println(minTemp)
		}
	}
}

func main() {
	var deptCount int

	_, err := fmt.Scanln(&deptCount)
	if err != nil {
		fmt.Println("Invalid departments number")

		return
	}

	for range deptCount {
		var staffCount int

		_, err := fmt.Scanln(&staffCount)
		if err != nil {
			fmt.Println("Invalid staff number")

			continue
		}

		processDept(staffCount)
	}
}
