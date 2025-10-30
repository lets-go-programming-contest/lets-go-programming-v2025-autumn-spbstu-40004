package main

import (
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

type TemperatureController struct {
	minBound int
	maxBound int
}

func createTempController() *TemperatureController {
	return &TemperatureController{
		minBound: minTemp,
		maxBound: maxTemp,
	}
}

func (tempcontr *TemperatureController) updateUpperLimit(requestedTemp int) {
	if tempcontr.minBound == -1 {
		return
	}

	if requestedTemp < tempcontr.minBound {
		tempcontr.minBound = -1

		return
	}

	if requestedTemp <= tempcontr.maxBound {
		tempcontr.maxBound = requestedTemp
	}
}

func (tempcontr *TemperatureController) updateLowerLimit(requestedTemp int) {
	if tempcontr.minBound == -1 {
		return
	}

	if requestedTemp > tempcontr.maxBound {
		tempcontr.minBound = -1

		return
	}

	if requestedTemp >= tempcontr.minBound {
		tempcontr.minBound = requestedTemp
	}
}

func (tempcontr *TemperatureController) adjustTemperatureRange(requestedTemp int, operation string) {
	switch operation {
	case ">=":
		tempcontr.updateLowerLimit(requestedTemp)
	case "<=":
		tempcontr.updateUpperLimit(requestedTemp)
	default:
		fmt.Println("Wrong input")

		return
	}
}

func (tempcontr *TemperatureController) determineComfortTemp() int {
	if tempcontr.minBound == -1 || tempcontr.minBound > tempcontr.maxBound {
		return -1
	}

	return tempcontr.minBound
}

func main() {
	var (
		numOfDeparts, numOfWorkers, currTemp int
		sign                                 string
	)

	_, err := fmt.Scanln(&numOfDeparts)
	if err != nil {
		fmt.Println("Wrong input")

		return
	}

	for i := 1; i <= numOfDeparts; i++ {
		_, err = fmt.Scanln(&numOfWorkers)
		if err != nil {
			fmt.Println("Wrong input")

			return
		}

		controller := createTempController()

		for j := 1; j <= numOfWorkers; j++ {
			_, err = fmt.Scanln(&sign, &currTemp)
			if err != nil {
				fmt.Println("Wrong input")

				return
			}

			controller.adjustTemperatureRange(currTemp, sign)
			optimalTemp := controller.determineComfortTemp()
			fmt.Println(optimalTemp)
		}
	}
}
