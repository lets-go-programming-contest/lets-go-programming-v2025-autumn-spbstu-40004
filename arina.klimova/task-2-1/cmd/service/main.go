package main

import (
	"fmt"
)

func findTemperature(peopleCount int) {

	minTemp := 15
	maxTemp := 30

	for range peopleCount {

		var operation string

		var personTemp int

		_, err := fmt.Scan(&operation, &personTemp)

		if err != nil {
			fmt.Println("Error reading operation or tempperature", err)

			return
		}

		switch operation {
		case ">=":
			if personTemp > minTemp {
				minTemp = personTemp
			}
		case "<=":
			if personTemp < maxTemp {
				maxTemp = personTemp
			}
		default:
			fmt.Println("invalid operation")

			return
		}

		if minTemp > maxTemp {
			fmt.Println("-1")
		} else {
			fmt.Println(minTemp)
		}
	}
}

func main() {
	var departmentCount int

	_, err := fmt.Scanln(&departmentCount)

	if err != nil {
		fmt.Println("Error reading department count:", err)

		return
	}

	for range departmentCount {
		var peopleCount int

		_, err := fmt.Scanln(&peopleCount)

		if err != nil {
			fmt.Println("Error reading people count:", err)

			return
		}

		findTemperature(peopleCount)
	}
}
