package main

import "fmt"

func adjustTemperature(lowTemp int, highTemp int, askingTemp int, operation string) (int, int) {
	if lowTemp == -1 && highTemp == -1 {
		return lowTemp, highTemp
	}

	switch operation {
	case ">=":
		if askingTemp > highTemp {
			lowTemp = -1
			highTemp = -1
		} else if lowTemp <= askingTemp && askingTemp <= highTemp {
			lowTemp = askingTemp
		}
	case "<=":
		if askingTemp < lowTemp {
			lowTemp = -1
			highTemp = -1
		} else if lowTemp <= askingTemp && askingTemp <= highTemp {
			highTemp = askingTemp
		}
	}

	return lowTemp, highTemp
}

func main() {
	var (
		departmentAmount, employeeAmount uint16
		askingTemp, lowTemp, highTemp    int
		operation                        string
	)

	_, err := fmt.Scanln(&departmentAmount)
	if err != nil {
		fmt.Println("Invalid department amount")

		return
	}

	for range departmentAmount {
		_, err = fmt.Scanln(&employeeAmount)
		if err != nil {
			fmt.Println("Invalid employee amount")

			return
		}

		lowTemp = 15
		highTemp = 30

		for range employeeAmount {
			_, err = fmt.Scanln(&operation, &askingTemp)
			if err != nil {
				fmt.Println("Invalid employee input")

				return
			}

			lowTemp, highTemp = adjustTemperature(lowTemp, highTemp, askingTemp, operation)
			fmt.Println(lowTemp)
		}
	}
}
