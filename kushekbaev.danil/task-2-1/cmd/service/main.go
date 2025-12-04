package main

import (
	"errors"
	"fmt"

	"github.com/Z-1337/task-2-1/internal/temperature"
)

var (
	ErrReadingDeps = errors.New("error while reading number of departments")
	ErrReadingEmpl = errors.New("error while reading number of employees")
	ErrReadingTemp = errors.New("error while reading desirable temperature")
	ErrInvalidOp   = errors.New("invalid operator")
)

const (
	minTemp      = 15
	maxTemp      = 30
	invalidValue = -1
)

func main() {
	var (
		departmentNumber int
		employeeNumber   int
		desirableTemp    int
		operator         string
	)

	_, err := fmt.Scan(&departmentNumber)
	if err != nil {
		fmt.Println(ErrReadingDeps)

		return
	}

	for range departmentNumber {
		_, err = fmt.Scan(&employeeNumber)
		if err != nil {
			fmt.Println(ErrReadingEmpl)

			return
		}

		temp := temperature.Temperature{
			LowerBound: minTemp,
			UpperBound: maxTemp,
		}

		for range employeeNumber {
			_, err = fmt.Scan(&operator, &desirableTemp)
			if err != nil {
				fmt.Println(ErrReadingTemp)

				return
			}

			if desirableTemp < minTemp || desirableTemp > maxTemp {
				fmt.Println(invalidValue)
			}

			finalTemp, err := temp.SetTemperature(operator, desirableTemp)
			if err != nil {
				fmt.Println(err)

				continue
			}

			fmt.Println(finalTemp)
		}
	}
}
