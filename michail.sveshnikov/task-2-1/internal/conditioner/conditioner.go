package conditioner

import (
	"errors"
)

var ErrConditioner = errors.New("invalid comparison sign")

type Conditioner struct {
	LowBound int
	UpBound  int
}

func (cond *Conditioner) OptimizeTemperature(compSign string, temperature int) (int, error) {
	const invalidTemperature int = -1

	switch compSign {
	case ">=":
		if temperature > cond.LowBound {
			cond.LowBound = temperature
		}
	case "<=":
		if temperature < cond.UpBound {
			cond.UpBound = temperature
		}
	default:
		return invalidTemperature, ErrConditioner
	}

	if cond.LowBound <= cond.UpBound {
		return cond.LowBound, nil
	} else {
		return invalidTemperature, nil
	}
}
