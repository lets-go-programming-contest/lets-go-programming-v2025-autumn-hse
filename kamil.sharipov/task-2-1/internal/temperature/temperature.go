package temperature

import "errors"

var ErrInvalidOperator = errors.New("Invalid operator")

const (
	minStartingTemp = 15
	maxStartingTemp = 30
)

const (
	opLessEqual    string = "<="
	opGreaterEqual string = ">="
)

func ParseOperator(s string) (string, error) {
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

func InitComfortTemperature() *ComfortTemperature {
	return &ComfortTemperature{
		min: minStartingTemp,
		max: maxStartingTemp,
	}
}

func (cr *ComfortTemperature) AddConstraint(op string, temp int) {
	switch op {
	case opLessEqual:
		if temp < cr.max {
			cr.max = temp
		}
	case opGreaterEqual:
		if temp > cr.min {
			cr.min = temp
		}
	}
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
