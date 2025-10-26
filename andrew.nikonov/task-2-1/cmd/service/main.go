package main

import (
	"errors"
	"fmt"

	TempManager "github.com/ysffmn/task-2-1/internal/TemperatureManager"
)

const (
	minAllowedTemp = 15
	maxAllowedTemp = 30
)

var (
	errDataFormat = errors.New("invalid temperature data")
	errDepNumber  = errors.New("invalid departments number")
	errEmplNumber = errors.New("invalid employee number")
)

func processDep(emplQuantity int) {
	depTemperature := TempManager.NewTemperatureManager(minAllowedTemp, maxAllowedTemp)

	for range emplQuantity {
		var (
			operator string
			temp     int
		)

		_, err := fmt.Scanln(&operator, &temp)
		if err != nil || temp < minAllowedTemp || temp > maxAllowedTemp {
			fmt.Println(errDataFormat)

			return
		}

		err = depTemperature.AdjustTemperature(operator, temp)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println(depTemperature.GetOptimalTemperature())
	}
}

func main() {
	var depQuantity, emplQuantity int

	_, err := fmt.Scanln(&depQuantity)
	if err != nil || depQuantity == 0 || depQuantity > 1000 {
		fmt.Println(errDepNumber)

		return
	}

	for range depQuantity {
		_, err := fmt.Scanln(&emplQuantity)
		if err != nil || emplQuantity == 0 || emplQuantity > 1000 {
			fmt.Println(errEmplNumber)

			return
		}

		processDep(emplQuantity)
	}
}
