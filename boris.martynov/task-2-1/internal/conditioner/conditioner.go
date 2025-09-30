package conditioner

import (
	"errors"
	"fmt"
)

var errFailedToScan = errors.New("invalid input")

const (
	MinTemp = 15
	MaxTemp = 30
)

type temperature struct {
	lowest  int
	highest int
}

func newTemperature(min, max int) *temperature {
	return &temperature{
		lowest:  min,
		highest: max,
	}
}

func (tr *temperature) update(greaterOrLess string, temp int) {
	switch greaterOrLess {
	case ">=":
		if temp > tr.lowest {
			tr.lowest = temp
		}
	case "<=":
		if temp < tr.highest {
			tr.highest = temp
		}
	}
}

func (tr *temperature) getResult() int {
	if tr.lowest <= tr.highest {
		return tr.lowest
	}
	return -1
}

func TemperatureWantedDepartment() {
	var departmentCapacity int

	if _, err := fmt.Scanln(&departmentCapacity); err != nil {
		fmt.Println(errFailedToScan)

		return
	}

	tr := newTemperature(MinTemp, MaxTemp)

	for range departmentCapacity {
		var (
			temperatureWantedByEmployee int
			greaterOrLess               string
		)

		if _, err := fmt.Scanln(&greaterOrLess, &temperatureWantedByEmployee); err != nil {
			fmt.Println(errFailedToScan)

			return
		}

		tr.update(greaterOrLess, temperatureWantedByEmployee)
		fmt.Println(tr.getResult())
	}
}
