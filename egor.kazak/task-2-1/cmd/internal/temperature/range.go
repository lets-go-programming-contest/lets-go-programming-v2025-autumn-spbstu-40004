package temperature

import "fmt"

const (
	DefaultMin = 15
	DefaultMax = 30
)

type Range struct {
	min int
	max int
}

func NewRange(min, max int) (*Range, error) {
	if min > max {
		return nil, fmt.Errorf("invalid range: min (%d) > max (%d)", min, max)
	}
	return &Range{min: min, max: max}, nil
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
		return fmt.Errorf("unsupported operator: %q", operator)
	}

	return nil
}

func (r *Range) Current() (int, bool) {
	if r.min <= r.max {
		return r.min, true
	}
	return 0, false
}
