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

func (temp *TemperatureRange) Update(targetTemp int, operation string) error {
	if temp.min == -1 && temp.max == -1 {
		return nil
	}

	switch operation {
	case ">=":
		if targetTemp > temp.max {
			temp.min = -1
			temp.max = -1
		} else if temp.min <= targetTemp && targetTemp <= temp.max {
			temp.min = targetTemp
		}
	case "<=":
		if targetTemp < temp.min {
			temp.min = -1
			temp.max = -1
		} else if temp.min <= targetTemp && targetTemp <= temp.max {
			temp.max = targetTemp
		}
	default:
		return ErrInvalidOperation
	}

	return nil
}

func (temp *TemperatureRange) GetMin() int {
	return temp.min
}

func (temp *TemperatureRange) GetMax() int {
	return temp.max
}
