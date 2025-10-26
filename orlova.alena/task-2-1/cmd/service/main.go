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

func newTemperatureController() *TemperatureController {
	return &TemperatureController{
		minBound: minTemp,
		maxBound: maxTemp,
	}
}

func (tempcontr *TemperatureController) changeMaxBound(currTemp int) {
	if tempcontr.minBound == -1 {
		return
	}

	if currTemp < tempcontr.minBound {
		tempcontr.minBound = -1

		return
	}

	if currTemp <= tempcontr.maxBound {
		tempcontr.maxBound = currTemp
	}
}

func (tempcontr *TemperatureController) changeMinBound(currTemp int) {
	if tempcontr.minBound == -1 {
		return
	}

	if currTemp > tempcontr.maxBound {
		tempcontr.minBound = -1

		return
	}

	if currTemp >= tempcontr.minBound {
		tempcontr.minBound = currTemp
	}
}

func (tempcontr *TemperatureController) findOptimalTemp(currTemp int, sign string) {
	switch sign {
	case ">=":
		tempcontr.changeMinBound(currTemp)
	case "<=":
		tempcontr.changeMaxBound(currTemp)
	default:
		fmt.Println("Wrong input")

		return
	}
}

func (tempcontr *TemperatureController) getTemperature() int {
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

		controller := newTemperatureController()

		for j := 1; j <= numOfWorkers; j++ {
			_, err = fmt.Scanln(&sign, &currTemp)
			if err != nil {
				fmt.Println("Wrong input")

				return
			}

			controller.findOptimalTemp(currTemp, sign)
			optimalTemp := controller.getTemperature()
			fmt.Println(optimalTemp)
		}
	}
}
