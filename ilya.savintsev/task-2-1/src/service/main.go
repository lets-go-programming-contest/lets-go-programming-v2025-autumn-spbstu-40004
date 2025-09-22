package main

import (
	"fmt"
)

func main() {
	var dep, emp int
	_, err := fmt.Scanln(&dep)
	if err != nil || dep < 1 || dep > 1000 {
		fmt.Println("Invalid number of departments")
		return
	}
	for i := 0; i < dep; i++ {
		_, err = fmt.Scanln(&emp)
		if err != nil || emp < 1 || emp > 1000 {
			fmt.Println("Invalid number of employees")
			return
		}
		lowTemp, highTemp := 15, 30
		tempCycle:
			for j := 0; j < emp; j++ {
				var op string
				var newTemp int
				_, err = fmt.Scanln(&op, &newTemp)
				if err != nil || newTemp < 15 || newTemp > 30 {
					fmt.Println("Invalid temperature format")
					return
				}
				switch op {
				case ">=":
					if highTemp < newTemp {
						fmt.Println(-1)
						break tempCycle
					}
					lowTemp = newTemp
				case "<=":
					if lowTemp > newTemp {
						fmt.Println(-1)
						break tempCycle
					}
					highTemp = newTemp
				default:
					fmt.Println("Unknown operation before temperature")
					return
				}
				fmt.Println(lowTemp)
			}
	}
}
