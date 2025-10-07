package conditioner

import (
	"fmt"
)

type temperature struct {
	lowest  int
	highest int
}

const (
	MinTemp = 15
	MaxTemp = 30
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
		return 0, fmt.Errorf("invalid temp range sign")
	}
	if tr.lowest <= tr.highest {
		return tr.lowest, nil
	}

	return -1, nil
}
