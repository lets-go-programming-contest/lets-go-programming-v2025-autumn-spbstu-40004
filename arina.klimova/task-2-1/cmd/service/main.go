package main

import (
	"fmt"
)

func main() {
	var daysCount, peopleCount int
	_, err := fmt.Scanln(&daysCount)

	if err != nil {
		fmt.Println("Error reading days count:", err)

		return
	}

	for i := 0; i < daysCount; i++ {
		_, err = fmt.Scanln(&peopleCount)

		if err != nil {
			fmt.Println("Error reading people count:", err)

			return
		}

		maxTemp := 30
		minTemp := 15

		for j := 0; j < peopleCount; j++ {
			var operation string
			var personTemp int

			_, err := fmt.Scan(&operation, &personTemp)

			if err != nil {
				fmt.Println("Error reading operation and temperature:", err)

				return
			}

			switch operation {
			case ">=":
				if personTemp >= minTemp && personTemp <= maxTemp {
					minTemp = personTemp
				}

				if maxTemp >= personTemp {
					fmt.Println(minTemp)
				} else {
					fmt.Println("-1")

					return
				}
			case "<=":
				if personTemp <= maxTemp && personTemp >= minTemp {
					maxTemp = personTemp
				}

				if minTemp <= personTemp {
					fmt.Println(minTemp)
				} else {
					fmt.Println("-1")

					return
				}
			default:
				fmt.Println("invalid operation")

				return
			}
		}
	}
}
