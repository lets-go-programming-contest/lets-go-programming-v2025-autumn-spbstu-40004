package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		maxTemp = 30
		minTemp = 15
		cntUnit = 0
	)

	_, err := fmt.Scan(&cntUnit)
	if err != nil || cntUnit < 1 || cntUnit > 1000 {
		fmt.Println("Incorrect input")

		return
	}
	for range cntUnit {
		var cntWorker int = 0
		_, err := fmt.Scan(&cntWorker)
		if err != nil || cntWorker < 1 || cntWorker > 1000 {
			fmt.Println("Incorrect input")

			return
		}

		temperature := make([]int, 0, maxTemp-minTemp+1)

		for filler := minTemp; filler < maxTemp+1; filler++ {
			temperature = append(temperature, filler)
		}

		for range cntWorker {
			var (
				comparator string
				tempValue  int
			)

			fmt.Scan(&comparator)
			_, err = fmt.Scan(&tempValue)
			if err != nil {
				fmt.Println("Incorrect input")

				return
			}

			switch {
			case strings.Compare(comparator, ">=") == 0:
				if tempValue > temperature[len(temperature)-1] {
					fmt.Println("-1")
				} else if tempValue > temperature[0] {
					temperature = temperature[tempValue-temperature[0]:]
				}

				fmt.Println(temperature[0])
			case strings.Compare(comparator, "<=") == 0:
				if tempValue < temperature[0] {
					fmt.Println("-1")
				} else if tempValue < temperature[len(temperature)-1] {
					temperature = temperature[:tempValue-temperature[0]+1]
				}

				fmt.Println(temperature[0])
			default:
				fmt.Println("Incorrect input")

				return
			}
		}
	}
}
