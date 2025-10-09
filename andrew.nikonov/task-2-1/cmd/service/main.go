package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

func main() {
	var depQuantity int

	_, err := fmt.Scanln(&depQuantity)
	if err != nil || depQuantity == 0 || depQuantity > 1000 {
		fmt.Println("Invalid departments number")
		return
	}

	for i := 0; i < depQuantity; i++ {
		var (
			emplQuantity int
			minDepTemp   = minTemp
			maxDepTemp   = maxTemp
		)

		_, err = fmt.Scanln(&emplQuantity)
		if err != nil || emplQuantity == 0 || emplQuantity > 1000 {
			fmt.Println("Invalid employee number")
			return
		}

		for j := 0; j < emplQuantity; j++ {
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
				return
			}
			fmt.Println(minDepTemp)
		}
	}
}
