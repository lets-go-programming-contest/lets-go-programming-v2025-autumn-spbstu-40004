package main

import "fmt"

func adjustTemperature(lowTemp int, highTemp int, askingTemp int, operation string) (int, int) {
	if lowTemp == -1 && highTemp == -1 {
		return -1, -1
	}

	switch operation {
	case ">=":
		if askingTemp > highTemp {
			return -1, -1
		}

		if askingTemp <= lowTemp {
			return lowTemp, highTemp
		}
		if lowTemp <= askingTemp && askingTemp <= highTemp {
			lowTemp = askingTemp
			return lowTemp, highTemp
		}
	case "<=":
		if askingTemp < lowTemp {
			return -1, -1
		}

		if askingTemp >= highTemp {
			return lowTemp, highTemp
		}

		if lowTemp <= askingTemp && askingTemp <= highTemp {
			highTemp = askingTemp
			return lowTemp, highTemp
		}
	}

	return -1, -1
}

func main() {
	var departmentAmount, employeeAmount uint16
	var askingTemp int
	var operation string

	_, err := fmt.Scanln(&departmentAmount)
	if err != nil {
		fmt.Println("Invalid department amount")
		return
	}

	for i := uint16(0); i < departmentAmount; i++ { // shrug
		_, err = fmt.Scanln(&employeeAmount)
		if err != nil {
			fmt.Println("Invalid employee amount")
			return
		}
		lowTemp := 15
		highTemp := 30

		for i := uint16(0); i < employeeAmount; i++ {

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
