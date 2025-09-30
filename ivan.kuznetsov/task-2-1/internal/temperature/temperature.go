package temperature

import "fmt"

type TemperatureRange struct {
	min int
	max int
}

func InitialRange() *TemperatureRange {
	return &TemperatureRange{
		min: 15,
		max: 30,
	}
}

func OptimalTemperature(sign string, value int, tr *TemperatureRange) *TemperatureRange {
	switch sign {
	case "<=":
		if value < tr.max {
			tr.max = value
		}
	case ">=":
		if value > tr.min {
			tr.min = value
		}
	default:
		fmt.Printf("Invalid comparison sign '%s'\nThe temperature range has not changed\n", sign)
	}
	return tr
}

func GetOptimalTemperature(tr *TemperatureRange) int {
	if tr.min > tr.max {
		return -1
	}
	return tr.min
}
