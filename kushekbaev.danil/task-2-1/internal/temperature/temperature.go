package temperature

import "errors"

var ErrInvalidOp = errors.New("invalid operator")

type Temperature struct {
	UpperBound int
	LowerBound int
}

func (temp *Temperature) SetTemperature(operator string, desirableTemp int) (int, error) {
	const invalidValue int = -1

	switch operator {
	case "<=":
		temp.UpperBound = ternaryInt(temp.UpperBound < desirableTemp, temp.UpperBound, desirableTemp)
	case ">=":
		temp.LowerBound = ternaryInt(temp.LowerBound > desirableTemp, temp.LowerBound, desirableTemp)
	default:
		return invalidValue, ErrInvalidOp
	}

	if temp.LowerBound <= temp.UpperBound {
		return temp.LowerBound, nil
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
