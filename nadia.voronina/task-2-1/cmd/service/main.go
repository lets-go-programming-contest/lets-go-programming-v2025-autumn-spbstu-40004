package main

import (
	"errors"
	"fmt"
)

const (
	MaxDegreeDefault = 30
	MinDegreeDefault = 15
)

var (
	ErrInvalidDepartments = errors.New("invalid number of departments")
	ErrInvalidEmployees   = errors.New("invalid number of employees")
	ErrInvalidSign        = errors.New("invalid sign")
	ErrInvalidDegree      = errors.New("invalid degree")
	ErrWrongSign          = errors.New("wrong sign has been added")
	ErrOutOfRange         = errors.New("degree out of range")
)

type DegreeRange struct {
	Max int
	Min int
}

func NewDegreeRange() *DegreeRange {
	return &DegreeRange{
		Max: MaxDegreeDefault,
		Min: MinDegreeDefault,
	}
}

func (dr *DegreeRange) ProcessCondition(degree int, sign string) (int, error) {
	if dr.Min == 0 && dr.Max == 0 {
		return -1, ErrOutOfRange
	}

	switch sign {
	case "<=":
		return dr.processLessEq(degree)
	case ">=":
		return dr.processGreaterEq(degree)
	default:
		return -1, ErrWrongSign
	}
}

func (dr *DegreeRange) processLessEq(degree int) (int, error) {
	switch {
	case dr.Max >= degree && dr.Min <= degree:
		dr.Max = degree
		return dr.Min, nil
	case dr.Max <= degree && dr.Min <= degree:
		return dr.Min, nil
	default:
		dr.Max = 0
		dr.Min = 0
		return -1, ErrOutOfRange
	}
}

func (dr *DegreeRange) processGreaterEq(degree int) (int, error) {
	switch {
	case dr.Min <= degree && dr.Max >= degree:
		dr.Min = degree
		return dr.Min, nil
	case dr.Min >= degree && dr.Max >= degree:
		return dr.Min, nil
	default:
		dr.Max = 0
		dr.Min = 0
		return -1, ErrOutOfRange
	}
}

func main() {
	var numberOfDepartments, numberOfEmployees int

	_, errN := fmt.Scanln(&numberOfDepartments)
	if errN != nil {
		fmt.Println(ErrInvalidDepartments)
		return
	}

	for range numberOfDepartments {
		_, errK := fmt.Scanln(&numberOfEmployees)
		if errK != nil {
			fmt.Println(ErrInvalidEmployees)
			return
		}

		dr := NewDegreeRange()

		for range numberOfEmployees {
			var sign string
			var degree int

			_, err1 := fmt.Scan(&sign)
			if err1 != nil {
				fmt.Println(ErrInvalidSign)
				return
			}

			_, err2 := fmt.Scanln(&degree)
			if err2 != nil {
				fmt.Println(ErrInvalidDegree)
				return
			}

			min, err := dr.ProcessCondition(degree, sign)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(min)
		}
	}
}
