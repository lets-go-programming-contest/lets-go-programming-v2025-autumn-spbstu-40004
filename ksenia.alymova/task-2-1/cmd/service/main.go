package main

import (
	"errors"
	"fmt"
	"strings"
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

	if tempValue < *lowBound && *lowBound != -1 {
		*lowBound = -1
	}

	if tempValue < *highBound {
		*highBound = tempValue
	}
}

func main() {
	var (
		cntUnit    = 0
		errorInput = errors.New("incorrect input")
	)

	_, err := fmt.Scan(&cntUnit)
	if err != nil || cntUnit < 1 || cntUnit > 1000 {
		fmt.Println(errorInput)

		return
	}

	for range cntUnit {
		const (
			minTemp = 15
			maxTemp = 30
		)
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
			var (
				comparator string
				tempValue  int
			)

			_, err = fmt.Scan(&comparator)
			if err != nil {
				fmt.Println(errorInput)

				return
			}

			_, err = fmt.Scan(&tempValue)
			if err != nil {
				fmt.Println(errorInput)

				return
			}

			switch {
			case strings.Compare(comparator, ">=") == 0:
				reduceLowBound(&lowBound, &highBound, tempValue)
			case strings.Compare(comparator, "<=") == 0:
				reduceHighBound(&lowBound, &highBound, tempValue)
			default:
				fmt.Println(errorInput)

				return
			}

			fmt.Println(lowBound)
		}
	}
}
