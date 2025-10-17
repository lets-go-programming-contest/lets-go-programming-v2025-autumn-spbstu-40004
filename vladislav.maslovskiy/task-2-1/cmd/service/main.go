package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidOperation = errors.New("invalid operation")
	ErrWrongDepartment  = errors.New("wrong department amount")
	ErrWrongEmployee    = errors.New("wrong employee amount")
	ErrWrongInput       = errors.New("wrong employee input")
	ErrInvalidTemp      = errors.New("invalid temperature")
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

func (tempContr *TemperatureController) changeMaxBound(currTemp int) {
	if tempContr.minBound == -1 {
		return
	}

	if currTemp < tempContr.minBound {
		tempContr.minBound = -1

		return
	}

	if currTemp <= tempContr.maxBound {
		tempContr.maxBound = currTemp
	}
}

func (tempContr *TemperatureController) changeMinBound(currTemp int) {
	if tempContr.minBound == -1 {
		return
	}

	if currTemp > tempContr.maxBound {
		tempContr.minBound = -1

		return
	}

	if currTemp >= tempContr.minBound {
		tempContr.minBound = currTemp
	}
}

func (tempContr *TemperatureController) findOptimalTemp(currTemp int, sign string) error {
	switch sign {
	case ">=":
		tempContr.changeMinBound(currTemp)
	case "<=":
		tempContr.changeMaxBound(currTemp)
	default:
		return ErrInvalidOperation
	}

	return nil
}

func (tempContr *TemperatureController) getTemperature() int {
	if tempContr.minBound == -1 || tempContr.minBound > tempContr.maxBound {
		return -1
	}

	return tempContr.minBound
}

func main() {
	var (
		numOfDeparts, numOfWorkers, currTemp int
		sign                                 string
	)

	_, err := fmt.Scanln(&numOfDeparts)
	if err != nil || numOfDeparts < 1 || numOfDeparts > 1000 {
		fmt.Println(ErrWrongDepartment)

		return
	}

	for range numOfDeparts {
		_, err = fmt.Scanln(&numOfWorkers)
		if err != nil || numOfWorkers < 1 || numOfWorkers > 1000 {
			fmt.Println(ErrWrongEmployee)

			return
		}

		controller := newTemperatureController()

		for range numOfWorkers {
			_, err = fmt.Scanln(&sign, &currTemp)
			if err != nil || currTemp < minTemp || currTemp > maxTemp {
				fmt.Println(ErrWrongInput)

				return
			}

			err = controller.findOptimalTemp(currTemp, sign)
			if err != nil {
				fmt.Println(err)

				return
			}

			optimalTemp := controller.getTemperature()
			fmt.Println(optimalTemp)
		}
	}
}

