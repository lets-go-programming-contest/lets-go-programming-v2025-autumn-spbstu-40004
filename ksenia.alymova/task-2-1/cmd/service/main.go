package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errorInput = errors.New("incorrect input")
)

const (
	minTemp = 15
	maxTemp = 30
)

func reduceLowBound(lowBound, highBound *int, tempValue int) {
	if *lowBound == -1 {
		return
	}

	if tempValue > *highBound {
		*lowBound = -1
	}

	if tempValue > *lowBound {
		*lowBound = tempValue
	}
}

func reduceHighBound(lowBound, highBound *int, tempValue int) {
	if *lowBound == -1 {
		return
	}

	if tempValue < *lowBound && *lowBound != -1 {
		*lowBound = -1
	}

	if tempValue < *highBound {
		*highBound = tempValue
	}
}

func changeTemperature(comparator string, tempValue int, lowBound, highBound *int) error {
	switch {
	case strings.Compare(comparator, ">=") == 0:
		reduceLowBound(lowBound, highBound, tempValue)
	case strings.Compare(comparator, "<=") == 0:
		reduceHighBound(lowBound, highBound, tempValue)
	default:
		return errorInput
	}

	return nil
}

func main() {
	var cntUnit int

	_, err := fmt.Scan(&cntUnit)
	if err != nil || cntUnit < 1 || cntUnit > 1000 {
		fmt.Println(errorInput)

		return
	}

	for range cntUnit {
		var (
			cntWorker = 0
			lowBound  = minTemp
			highBound = maxTemp
		)

		_, err = fmt.Scan(&cntWorker)
		if err != nil || cntWorker < 1 || cntWorker > 1000 {
			fmt.Println(errorInput)

			return
		}

		for range cntWorker {
			var comparator string
			var tempValue int

			_, err = fmt.Scan(&comparator)
			if err != nil {
				fmt.Println(errorInput)

				return
			}

			_, err = fmt.Scan(&tempValue)
			if err != nil || tempValue < minTemp || tempValue > maxTemp {
				fmt.Println(errorInput)

				return
			}

			err = changeTemperature(comparator, tempValue, &lowBound, &highBound)
			if err != nil {
				fmt.Println(err)

				return
			}

			fmt.Println(lowBound)
		}
	}
}
