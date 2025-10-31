package temperature

import (
	"errors"
)

var (
	ErrInvalidOp = errors.New("invalid operator")
)

type Temperature struct {
	UpperBound int
	LowerBound int
}

func (temp *Temperature) setTemperature(operator string, desirableTemp int) {
	const invalidValue int = -1;

	switch operator {
	case "<=":
		temp.UpperBound = ternaryInt(UpperBound < desirableTemp, UpperBound, desirableTemp)
	case ">=":
		temp.LowerBound = ternaryInt(LowerBound > desirableTemp, LowerBound, desirableTemp)
	default:
		return invalidValue, ErrInvalidOp
	}

	if lowerBound <= upperBound {
		return cond.LowerBound, nil
	} else {
		return invalidValue, nil
	}
}

func ternaryInt(condition bool, trueValue int, falseValue int) int {
	if condition {
		return trueValue
	}

	return falseValue
}