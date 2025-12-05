package main

import (
	"errors"
	"fmt"
)

const (
	defaultMaxDegree = 30
	defaultMinDegree = 15
	invalidState     = 0
	noSolution       = -1
	signLessEq       = "<="
	signGreaterEq    = ">="
)

const (
	msgInvalidDepartments = "Invalid number of departments"
	msgInvalidEmployees   = "Invalid number of employees"
	msgInvalidSign        = "Invalid sign"
	msgInvalidDegree      = "Invalid degree"
	msgWrongSign          = "Wrong sign has been added"
)

var (
	errInvalidSign = errors.New("wrong sign has been added")
	errNoSolution  = errors.New("no solution")
)

type TemperatureRange struct {
	maxDegree int
	minDegree int
}

func NewTemperatureRange() *TemperatureRange {
	return &TemperatureRange{
		maxDegree: defaultMaxDegree,
		minDegree: defaultMinDegree,
	}
}

func (t *TemperatureRange) ProcessCondition(degree int, sign string) (int, error) {
	if t.minDegree == invalidState && t.maxDegree == invalidState {
		return noSolution, errNoSolution
	}

	switch sign {
	case signLessEq:
		return t.processLessEq(degree)
	case signGreaterEq:
		return t.processGreaterEq(degree)
	default:
		return 0, errInvalidSign
	}
}

func (t *TemperatureRange) processLessEq(degree int) (int, error) {
	switch {
	case t.maxDegree >= degree && t.minDegree <= degree:
		t.maxDegree = degree

		return t.minDegree, nil
	case t.maxDegree <= degree && t.minDegree <= degree:
		return t.minDegree, nil
	default:
		t.maxDegree = invalidState
		t.minDegree = invalidState

		return noSolution, errNoSolution
	}
}

func (t *TemperatureRange) processGreaterEq(degree int) (int, error) {
	switch {
	case t.minDegree <= degree && t.maxDegree >= degree:
		t.minDegree = degree

		return t.minDegree, nil
	case t.minDegree >= degree && t.maxDegree >= degree:
		return t.minDegree, nil
	default:
		t.maxDegree = invalidState
		t.minDegree = invalidState

		return noSolution, errNoSolution
	}
}

func main() {
	var numberOfDepartments, numberOfEmployees int

	_, errN := fmt.Scanln(&numberOfDepartments)
	if errN != nil {
		fmt.Println(msgInvalidDepartments)

		return
	}

	for range numberOfDepartments {
		_, errK := fmt.Scanln(&numberOfEmployees)
		if errK != nil {
			fmt.Println(msgInvalidEmployees)

			return
		}

		tempRange := NewTemperatureRange()

		for range numberOfEmployees {
			var sign string

			var degree int

			_, err1 := fmt.Scan(&sign)
			if err1 != nil {
				fmt.Println(msgInvalidSign)

				return
			}

			_, err2 := fmt.Scanln(&degree)
			if err2 != nil {
				fmt.Println(msgInvalidDegree)

				return
			}

			result, err := tempRange.ProcessCondition(degree, sign)
			if err != nil {
				if errors.Is(err, errInvalidSign) {
					fmt.Println(msgWrongSign)

					return
				}
			}

			fmt.Println(result)
		}
	}
}
