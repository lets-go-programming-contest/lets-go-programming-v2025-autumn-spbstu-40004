package temperature

import (
    "errors"
    "fmt"
)

const (
	DefaultMin = 15
	DefaultMax = 30
)

type Range struct {
	min int
	max int
}

var (
	ErrInvalidRange     = errors.New("invalid range: min > max")
	ErrUnsupportedOp    = errors.New("unsupported operator")
)

func NewRange(minVal, maxVal int) (*Range, error) {
	if minVal > maxVal {
		return nil, fmt.Errorf("%w: %d > %d", ErrInvalidRange, minVal, maxVal)
	}

	return &Range{min: minVal, max: maxVal}, nil
}

func NewDefaultRange() *Range {
	return &Range{min: DefaultMin, max: DefaultMax}
}

func (r *Range) ApplyConstraint(operator string, temperature int) error {
	switch operator {
	case ">=":
		if temperature > r.min {
			r.min = temperature
		}
	case "<=":
		if temperature < r.max {
			r.max = temperature
		}
	default:
		return fmt.Errorf("%w: %q", ErrUnsupportedOp, operator)
	}

	return nil
}

func (r *Range) Current() (int, bool) {
	if r.min <= r.max {

		return r.min, true
	}

	return 0, false
}
