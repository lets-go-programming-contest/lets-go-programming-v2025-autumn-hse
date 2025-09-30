package temperature

import "fmt"

type TemperatureRange struct {
	Min int
	Max int
}

func OptimalTemperature(sign string, value int, tr *TemperatureRange) *TemperatureRange {
	switch sign {
	case "<=":
		if value < tr.Max {
			tr.Max = value
		}
	case ">=":
		if value > tr.Min {
			tr.Min = value
		}
	default:
		fmt.Printf("Invalid comparison sign '%s'\nThe temperature range has not changed\n", sign)
	}

	return tr
}

func GetOptimalTemperature(tr *TemperatureRange) int {
	if tr.Min > tr.Max {
		return -1
	}

	return tr.Min
}
