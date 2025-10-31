package temperature

import "errors"

var (
	ErrInvalidOperation = errors.New("invalid operation")
	ErrInvalidInput     = errors.New("invalid employee input")
)

const (
	MinTemperature = 15
	MaxTemperature = 30
)

type TemperatureRange struct {
	min int
	max int
}

func NewTemperatureRange(minTemp, maxTemp int) *TemperatureRange {
	return &TemperatureRange{
		min: minTemp,
		max: maxTemp,
	}
}

func (tr *TemperatureRange) Update(targetTemp int, operation string) error {
	if tr.min == -1 && tr.max == -1 {
		return nil
	}

	switch operation {
	case ">=":
		if targetTemp > tr.max {
			tr.min = -1
			tr.max = -1
		} else if tr.min <= targetTemp && targetTemp <= tr.max {
			tr.min = targetTemp
		}
	case "<=":
		if targetTemp < tr.min {
			tr.min = -1
			tr.max = -1
		} else if tr.min <= targetTemp && targetTemp <= tr.max {
			tr.max = targetTemp
		}
	default:
		return ErrInvalidOperation
	}

	return nil
}

func (tr *TemperatureRange) GetMin() int {
	return tr.min
}

func (tr *TemperatureRange) GetMax() int {
	return tr.max
}
