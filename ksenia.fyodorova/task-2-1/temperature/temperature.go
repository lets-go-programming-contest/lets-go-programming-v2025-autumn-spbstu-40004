package temperature

import "errors"

var (
	ErrInvalidRange = errors.New("invalid temperature range")
	ErrNoSolution   = errors.New("no temperature satisfies all constraints")
)

type TemperatureRange struct {
	Min int
	Max int
}

func NewTemperatureRange() *TemperatureRange {
	return &TemperatureRange{
		Min: 15,
		Max: 30,
	}
}

func (tr *TemperatureRange) Update(sign string, temperature int) error {
	switch sign {
	case ">=":
		if temperature > tr.Min {
			tr.Min = temperature
		}
	case "<=":
		if temperature < tr.Max {
			tr.Max = temperature
		}
	default:
		return ErrInvalidRange
	}

	if tr.Min > tr.Max {
		return ErrNoSolution
	}

	return nil
}

func (tr *TemperatureRange) GetComfortableTemp() (int, error) {
	if tr.Min > tr.Max {
		return -1, ErrNoSolution
	}
	return tr.Min, nil
}
