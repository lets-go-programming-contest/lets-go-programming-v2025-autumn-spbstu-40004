package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&K)
		var maxTemp int = 30
		var minTemp int = 15
		for j := 0; j < K; j++ {
			var op string
			var personTemp int
			fmt.Scan(&op, &personTemp)
			if op == ">=" {
				if personTemp >= minTemp && personTemp <= maxTemp {
					minTemp = personTemp
				}
			} else if op == "<=" {
				if personTemp <= maxTemp && personTemp >= minTemp {
					maxTemp = personTemp
				}
			}
			if maxTemp >= personTemp && minTemp <= personTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println("-1")
				return
			}
		}
	}
}
