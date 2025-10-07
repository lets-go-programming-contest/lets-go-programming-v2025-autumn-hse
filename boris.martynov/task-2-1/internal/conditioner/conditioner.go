package conditioner

import "errors"

type temperature struct {
	lowest  int
	highest int
}

const (
	MinTemp = 15
	MaxTemp = 30
)

var (
	ErrInvalidTempSign = errors.New("invalid temp range sign")
)

func NewTemperature(low, high int) *temperature {
	return &temperature{
		lowest:  low,
		highest: high,
	}
}

func (tr *temperature) TempWantedByEmployee(greaterOrLess string, temp int) (int, error) {
	switch greaterOrLess {
	case ">=":
		if temp > tr.lowest {
			tr.lowest = temp
		}
	case "<=":
		if temp < tr.highest {
			tr.highest = temp
		}
	default:
		return 0, ErrInvalidTempSign
	}

	if tr.lowest <= tr.highest {
		return tr.lowest, nil
	}

	return -1, nil
}
