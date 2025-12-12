package temperature

import "errors"

var (
	ErrInvalidRange = errors.New("invalid temperature range")
	ErrNoSolution   = errors.New("no temperature satisfies all constraints")
)

const (
	DefaultMinTemp = 15
	DefaultMaxTemp = 30
)

type TemperatureRange struct {
	min int
	max int
}

func NewTemperatureRange() *TemperatureRange {
	return &TemperatureRange{
		min: DefaultMinTemp,
		max: DefaultMaxTemp,
	}
}

func (tr *TemperatureRange) Update(sign string, temperature int) error {
	switch sign {
	case ">=":
		if temperature > tr.min {
			tr.min = temperature
		}
	case "<=":
		if temperature < tr.max {
			tr.max = temperature
		}
	default:
		return ErrInvalidRange
	}

	if tr.min > tr.max {
		return ErrNoSolution
	}

	return nil
}

func (tr *TemperatureRange) GetComfortableTemp() (int, error) {
	if tr.min > tr.max {
		return -1, ErrNoSolution
	}

	return tr.min, nil
}

func (tr *TemperatureRange) GetMin() int {
	return tr.min
}

func (tr *TemperatureRange) GetMax() int {
	return tr.max
}
