package main

import (
	"fmt"
	"strings"
)

func reduceLowBound(lowBound, highBound, tempValue int) (int, int) {
	if lowBound == -1 {
		return lowBound, highBound
	}

	if tempValue > highBound {
		lowBound = -1

		return lowBound, highBound
	}

	if tempValue > lowBound {
		lowBound = tempValue
	}

	return lowBound, highBound
}

func reduceHighBound(lowBound, highBound, tempValue int) (int, int) {
	if lowBound == -1 {
		return lowBound, highBound
	}

	if tempValue < lowBound && lowBound != -1 {
		lowBound = -1

		return lowBound, highBound
	}

	if tempValue < highBound {
		highBound = tempValue
	}

	return lowBound, highBound
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
		var (
			cntWorker = 0
			lowBound  = 15
			highBound = 30
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
				lowBound, highBound = reduceLowBound(lowBound, highBound, tempValue)
				res = append(res, lowBound)
			case strings.Compare(comparator, "<=") == 0:
				lowBound, highBound = reduceHighBound(lowBound, highBound, tempValue)
				res = append(res, lowBound)
			default:
				fmt.Println("Incorrect input")

				return
			}
		}
	}

	for indexRes := range res {
		fmt.Println(res[indexRes])
	}
}
