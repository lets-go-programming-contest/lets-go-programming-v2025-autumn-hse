package conditioner

import (
	"fmt"
)

const (
	MinTemp = 15
	MaxTemp = 30
)

type temperature struct {
	lowest  int
	highest int
}

func newTemperature(low, high int) *temperature {
	return &temperature{
		lowest:  low,
		highest: high,
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
		fmt.Println("when scanning capacity of department: ", err)

		return
	}

	temperatureRange := newTemperature(MinTemp, MaxTemp)

	for range departmentCapacity {
		var (
			temperatureWantedByEmployee int
			greaterOrLess               string
		)

		if _, err := fmt.Scanln(&greaterOrLess, &temperatureWantedByEmployee); err != nil {
			fmt.Println("when scanning temperature wanted by employee: ", err)

			return
		}

		temperatureRange.update(greaterOrLess, temperatureWantedByEmployee)
		fmt.Println(temperatureRange.getResult())
	}
}
