package main

import (
	"fmt"
	"log"
)

func main() {
	var count int

	_, err := fmt.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	for range count {
		var constraints int
		_, err := fmt.Scan(&constraints)
		if err != nil {
			log.Fatal(err)
		}

		minTemp := 15
		maxTemp := 30

		for range constraints {
			var operator string
			
			var temperature int

			_, err := fmt.Scan(&operator, &temperature)
			if err != nil {
				log.Fatal(err)
			}

			switch operator {
			case ">=":
				if temperature > minTemp {
					minTemp = temperature
				}
			case "<=":
				if temperature < maxTemp {
					maxTemp = temperature
				}
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
