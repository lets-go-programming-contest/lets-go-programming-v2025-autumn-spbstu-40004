package TemperatureManager

import (
	"errors"
)

var (
	ErrOperator = errors.New("invalid operator")
)

type TemperatureManager struct {
	minTemp int
	maxTemp int
}

func NewTemperatureManager(minT int, maxT int) *TemperatureManager {
	return &TemperatureManager{minTemp: minT, maxTemp: maxT}
}

func (t *TemperatureManager) AdjustTemperature(operator string, val int) error {
	switch operator {
	case "<=":
		t.maxTemp = min(t.maxTemp, val)

		return nil
	case ">=":
		t.minTemp = max(t.minTemp, val)

		return nil
	default:
		return ErrOperator
	}
}

func (t *TemperatureManager) GetOptimalTemperature() int {
	if t.minTemp > t.maxTemp {
		return -1
	} else {
		return t.minTemp
	}
}
