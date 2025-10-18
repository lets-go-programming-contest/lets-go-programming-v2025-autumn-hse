package temperature

import "errors"

type SignType string

var ErrInvalidOperator = errors.New("invalid operator")

const (
	minStartingTemp = 15
	maxStartingTemp = 30
)

const (
	opLessEqual    SignType = "<="
	opGreaterEqual SignType = ">="
)

func ParseOperator(s string) (SignType, error) {
	switch s {
	case "<=":
		return opLessEqual, nil
	case ">=":
		return opGreaterEqual, nil
	default:
		return "", ErrInvalidOperator
	}
}

type ComfortTemperature struct {
	max int
	min int
}

func NewComfortTemperature() ComfortTemperature {
	return ComfortTemperature{
		min: minStartingTemp,
		max: maxStartingTemp,
	}
}

func (cr *ComfortTemperature) AddConstraint(operator SignType, temp int) error {
	switch operator {
	case opLessEqual:
		cr.max = min(cr.max, temp)
	case opGreaterEqual:
		cr.min = max(cr.min, temp)
	default:
		return ErrInvalidOperator
	}

	return nil
}

func (cr *ComfortTemperature) isValid() bool {
	return cr.min <= cr.max
}

func (cr *ComfortTemperature) Result() int {
	if cr.isValid() {
		return cr.min
	}

	return -1
}
