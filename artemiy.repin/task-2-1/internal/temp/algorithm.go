package temp

import "errors"

const (
	MinBound = 15
	MaxBound = 30
)

type Temperature struct {
	LeftBound  int
	RightBound int
}

var ErrUnsupportedOperator = errors.New("unsupported operator")

func NewTemperature(left, right int) Temperature {
	return Temperature{
		LeftBound:  left,
		RightBound: right,
	}
}

func (t *Temperature) UpdateInterval(operator string, value int) error {
	switch operator {
	case "<=":
		if value < t.RightBound {
			t.RightBound = value
		}
	case ">=":
		if value > t.LeftBound {
			t.LeftBound = value
		}
	default:
		return ErrUnsupportedOperator
	}

	return nil
}

func (t *Temperature) GetOptimal() int {
	if t.LeftBound <= t.RightBound {
		return t.LeftBound
	}

	return -1
}
