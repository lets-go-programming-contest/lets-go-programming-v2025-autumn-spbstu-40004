package temperature

import (
	"errors"
)

const (
	minimalTemp = 15
	maximalTemp = 30
)

var (
	ErrInvalidOperation = errors.New("invalid operation")
	ErrInvalidTempRange = errors.New("temperature range is invalid")
)

type TemperatureController struct {
	minTemp int
	maxTemp int
}

func NewTemperatureController() *TemperatureController {
	return &TemperatureController{
		minTemp: minimalTemp,
		maxTemp: maximalTemp,
	}
}

func (tc *TemperatureController) UpdateTemperature(operation string, personTemp int) error {
	switch operation {
	case ">=":
		if personTemp > tc.minTemp {
			tc.minTemp = personTemp
		}
	case "<=":
		if personTemp < tc.maxTemp {
			tc.maxTemp = personTemp
		}
	default:
		return ErrInvalidOperation
	}

	return nil
}

func (tc *TemperatureController) GetCurrentTemperature() (int, error) {
	if tc.minTemp > tc.maxTemp {
		return -1, ErrInvalidTempRange
	}

	return tc.minTemp, nil
}
