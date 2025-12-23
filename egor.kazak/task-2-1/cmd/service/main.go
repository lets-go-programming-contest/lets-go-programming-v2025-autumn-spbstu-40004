package main

import (
	"fmt"
	"log"

	"task-2-1/internal/temperature"
)

func main() {
	var count int
	if _, err := fmt.Scan(&count); err != nil {
		log.Fatal(fmt.Errorf("failed to read count: %w", err))
	}

	for range count {
		var constraints int
		if _, err := fmt.Scan(&constraints); err != nil {
			log.Fatal(fmt.Errorf("failed to read constraints count: %w", err))
		}

		tempRange := temperature.NewDefaultRange()

		for range constraints {
			var (
				operator    string
				temperature int
			)
			if _, err := fmt.Scan(&operator, &temperature); err != nil {
				log.Fatal(fmt.Errorf("failed to read constraint: %w", err))
			}

			if err := tempRange.ApplyConstraint(operator, temperature); err != nil {
				log.Fatal(fmt.Errorf("invalid constraint %q %d: %w", operator, temperature, err))
			}

			if value, ok := tempRange.Current(); ok {
				fmt.Println(value)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
