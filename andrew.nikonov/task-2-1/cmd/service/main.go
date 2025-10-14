package main

import (
	"errors"
	"fmt"
)

const (
	minAllowedTemp = 15
	maxAllowedTemp = 30
)

type Temperature struct {
	minTemp int
	maxTemp int
}

func (t *Temperature) adjustTemperature(operator string, val int) error {
	switch operator {
	case "<=":
		t.maxTemp = min(t.maxTemp, val)

		return nil
	case ">=":
		t.minTemp = max(t.minTemp, val)

		return nil
	default:
		return errOperator
	}
}

func (t *Temperature) getOptimalTemperature() int {
	if t.minTemp > t.maxTemp {
		return -1
	} else {
		return t.minTemp
	}
}

var (
	errDataFormat = errors.New("invalid temperature data")
	errOperator   = errors.New("invalid operator")
	errDepNumber  = errors.New("invalid departments number")
	errEmplNumber = errors.New("invalid employee number")
)

func processDep(emplQuantity int) {
	depTemperature := Temperature{minTemp: minAllowedTemp, maxTemp: maxAllowedTemp}
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

		err = depTemperature.adjustTemperature(operator, temp)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println(depTemperature.getOptimalTemperature())
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
