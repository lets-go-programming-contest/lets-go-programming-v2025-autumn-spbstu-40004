package main

import (
	"errors"
)

var errOperation = errors.New("Invalid operation")

func updateTemperatureRange(minTemp int, maxTemp int, targetTemp int, operation string) (int, int, error) {
	if minTemp == -1 && maxTemp == -1 {
		return minTemp, maxTemp, nil
	}

	switch operation {
	case ">=":
		if targetTemp > maxTemp {
			minTemp = -1
			maxTemp = -1
		} else if minTemp <= targetTemp && targetTemp <= maxTemp {
			minTemp = targetTemp
		}
	case "<=":
		if targetTemp < minTemp {
			minTemp = -1
			maxTemp = -1
		} else if minTemp <= targetTemp && targetTemp <= maxTemp {
			maxTemp = targetTemp
		}
	default:
		return minTemp, maxTemp, errOperation
	}

	return minTemp, maxTemp, nil
}