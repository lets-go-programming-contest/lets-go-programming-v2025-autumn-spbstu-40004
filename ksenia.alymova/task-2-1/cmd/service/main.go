package main

import (
	"errors"
	"fmt"
	"strings"
)

var errInput = errors.New("incorrect input")

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

		return
	}

	if tempValue > *lowBound {
		*lowBound = tempValue
	}
}

func reduceHighBound(lowBound, highBound *int, tempValue int) {
	if *lowBound == -1 {
		return
	}

	if tempValue < *lowBound {
		*lowBound = -1

		return
	}

	if tempValue < *highBound {
		*highBound = tempValue
	}
}

func changeTemperature(lowBound, highBound *int) error {
	var (
		comparator string
		tempValue  int
	)

	_, err := fmt.Scanln(&comparator, &tempValue)
	if err != nil || tempValue < minTemp || tempValue > maxTemp {
		return errInput
	}

	switch {
	case strings.Compare(comparator, ">=") == 0:
		reduceLowBound(lowBound, highBound, tempValue)
	case strings.Compare(comparator, "<=") == 0:
		reduceHighBound(lowBound, highBound, tempValue)
	default:
		return errInput
	}

	return nil
}

func main() {
	var cntUnit int

	_, err := fmt.Scanln(&cntUnit)
	if err != nil || cntUnit < 1 || cntUnit > 1000 {
		fmt.Println(errInput)

		return
	}

	for range cntUnit {
		var (
			cntWorker int
			lowBound  = minTemp
			highBound = maxTemp
		)

		_, err = fmt.Scanln(&cntWorker)
		if err != nil || cntWorker < 1 || cntWorker > 1000 {
			fmt.Println(errInput)

			return
		}

		for range cntWorker {
			err = changeTemperature(&lowBound, &highBound)
			if err != nil {
				fmt.Println(errInput)

				return
			}

			fmt.Println(lowBound)
		}
	}
}
