package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

func processDep(emplQuantity int) {
	var (
		minDepTemp = minTemp
		maxDepTemp = maxTemp
	)

	for range emplQuantity {
		var (
			operator string
			temp     int
		)

		_, err := fmt.Scanln(&operator, &temp)
		if err != nil || temp < minTemp || temp > maxTemp {
			fmt.Println("Invalid temperature data")

			return
		}

		switch operator {
		case "<=":
			maxDepTemp = min(maxDepTemp, temp)
		case ">=":
			minDepTemp = max(minDepTemp, temp)
		default:
			fmt.Println("Invalid operator")

			return
		}

		if minDepTemp > maxDepTemp {
			fmt.Println("-1")
		} else {
			fmt.Println(minDepTemp)
		}
	}
}

func main() {
	var depQuantity, emplQuantity int

	_, err := fmt.Scanln(&depQuantity)
	if err != nil || depQuantity == 0 || depQuantity > 1000 {
		fmt.Println("Invalid departments number")

		return
	}

	for range depQuantity {
		_, err := fmt.Scanln(&emplQuantity)
		if err != nil || emplQuantity == 0 || emplQuantity > 1000 {
			fmt.Println("Invalid employee number")

			return
		}

		processDep(emplQuantity)
	}
}
