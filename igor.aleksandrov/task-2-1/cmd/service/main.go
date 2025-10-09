package main

import "fmt"

const (
	minTemperature     = 15
	maxTemperature     = 30
	invalidTemperature = -1
)

func printOptimalTemperature(employeesCount int) {
	minTemperatureValue := minTemperature
	maxTemperatureValue := maxTemperature

	for range employeesCount {
		var comparisonSign string

		_, err := fmt.Scan(&comparisonSign)
		if err != nil {
			fmt.Println("Invalid comparison sign for temperature!")

			continue
		}

		var temperature int

		_, err = fmt.Scan(&temperature)
		if err != nil {
			fmt.Println("Invalid temperature value!")

			continue
		}

		if temperature < minTemperature || temperature > maxTemperature {
			fmt.Println("Unsupported temperature value!")

			continue
		}

		switch comparisonSign {
		case ">=":
			if temperature > minTemperatureValue {
				minTemperatureValue = temperature
			}
		case "<=":
			if temperature < maxTemperatureValue {
				maxTemperatureValue = temperature
			}
		default:
			fmt.Println("Unsupported comparison sign for temperature!")

			continue
		}

		if minTemperatureValue > maxTemperatureValue {
			fmt.Println(invalidTemperature)
		} else {
			fmt.Println(minTemperatureValue)
		}
	}
}

func main() {
	var departmentsCount int

	_, err := fmt.Scanln(&departmentsCount)
	if err != nil || departmentsCount < 1 || departmentsCount > 1000 {
		fmt.Println("Invalid departments count!")

		return
	}

	for range departmentsCount {
		var employeesCount int

		_, err = fmt.Scanln(&employeesCount)
		if err != nil || employeesCount < 1 || employeesCount > 1000 {
			fmt.Println("Invalid employees count!")

			continue
		}

		printOptimalTemperature(employeesCount)
	}
}
