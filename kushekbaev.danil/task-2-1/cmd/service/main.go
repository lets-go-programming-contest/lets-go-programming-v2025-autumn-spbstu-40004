package main

import (
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
	invalidValue = -1
)

func main() {
	var (
		departmentNumber int
		employeeNumber   int
		desireableTemp   int
		operator         string
	)

	_, err := fmt.Scan(&departmentNumber)
	if err != nil {
		fmt.Println("Error while reading number of departments")

		return
	}

	for range departmentNumber {
		_, err = fmt.Scan(&employeeNumber)
		if err != nil {
			fmt.Println("Error while reading number of employees")

			return
		}

		lowerBound := minTemp
		upperBound := maxTemp

		for range employeeNumber {
			_, err = fmt.Scan(&operator)
			if err != nil {
				fmt.Println("Error while reading operator")

				return
			}

			_, err = fmt.Scan(&desireableTemp)
			if err != nil {
				fmt.Println("Error while reading desireable temperature")

				return
			}

			if desireableTemp < minTemp || desireableTemp > maxTemp {
				fmt.Println(invalidValue)
			}

			switch operator {
			case "<=":
				upperBound = ternaryInt(upperBound < desireableTemp, upperBound, desireableTemp)
			case ">=":
				lowerBound = ternaryInt(lowerBound > desireableTemp, lowerBound, desireableTemp)
			default:
				fmt.Println("Invalid operator")
			}

			if (lowerBound <= upperBound) {
				fmt.Println(lowerBound)
			} else {
				fmt.Println(invalidValue)
			}
		}
	}
}

func ternaryInt(condition bool, trueValue int, falseValue int) int {
	if condition {
		return trueValue
	}
	return falseValue
}
