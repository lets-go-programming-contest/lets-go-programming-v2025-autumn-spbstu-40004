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
