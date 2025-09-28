package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		max_temp int = 30
		min_temp int = 15
		cnt_unit int = 0
	)
	_, err := fmt.Scan(&cnt_unit)
	if err != nil || cnt_unit < 1 || cnt_unit > 1000 {
		fmt.Println("Incorrect input")
		return
	}
	for index_unit := 0; index_unit < cnt_unit; index_unit++ {
		var cnt_worker int = 0
		_, err := fmt.Scan(&cnt_worker)
		if err != nil || cnt_worker < 1 || cnt_worker > 1000 {
			fmt.Println("Incorrect input")
			return
		}
		temperature := make([]int, 0, max_temp-min_temp+1)
		for filler := min_temp; filler < max_temp+1; filler++ {
			temperature = append(temperature, filler)
		}
		for index_worker := 0; index_worker < cnt_worker; index_worker++ {
			var (
				comparator string
				temp_value int
			)
			_, err = fmt.Scan(&comparator, &temp_value)
			if err != nil {
				fmt.Println("Incorrect input")
				return
			}
			if strings.Compare(comparator, ">=") == 0 {
				if temp_value > temperature[len(temperature)-1] {
					fmt.Println("-1")
				} else if temp_value > temperature[0] {
					temperature = temperature[temp_value-temperature[0]:]
				}
			} else if strings.Compare(comparator, "<=") == 0 {
				if temp_value < temperature[0] {
					fmt.Println("-1")
				} else if temp_value < temperature[len(temperature)-1] {
					temperature = temperature[:temp_value-temperature[0]+1]
				}
			} else {
				fmt.Println("Incorrect input")
				return
			}
			fmt.Println(temperature[0])
			//fmt.Println(temperature)
		}

	}

}
