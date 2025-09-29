package main

import (
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
		cntUnit = 0
		res     = make([]int, 0)
	)

	_, err := fmt.Scan(&cntUnit)
	if err != nil || cntUnit < 1 || cntUnit > 1000 {
		fmt.Println("Incorrect input")

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
			fmt.Println("Incorrect input")

			return
		}

		for range cntWorker {
			var (
				comparator string
				tempValue  int
			)

			_, err = fmt.Scan(&comparator)
			if err != nil {
				fmt.Println("Incorrect input")

				return
			}

			_, err = fmt.Scan(&tempValue)
			if err != nil {
				fmt.Println("Incorrect input")

				return
			}

			switch {
			case strings.Compare(comparator, ">=") == 0:
				reduceLowBound(&lowBound, &highBound, tempValue)
			case strings.Compare(comparator, "<=") == 0:
				reduceHighBound(&lowBound, &highBound, tempValue)
			default:
				fmt.Println("Incorrect input")

				return
			}
			res = append(res, lowBound)
		}
	}

	for indexRes := range res {
		fmt.Println(res[indexRes])
	}
}
