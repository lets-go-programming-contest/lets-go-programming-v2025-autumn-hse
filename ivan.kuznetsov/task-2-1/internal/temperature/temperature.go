package temperature

import "fmt"

type TemperatureRange struct {
	Min int
	Max int
}

func OptimalTemperature(sign string, value int, temperature *TemperatureRange) *TemperatureRange {
	switch sign {
	case "<=":
		if value < temperature.Max {
			temperature.Max = value
		}
	case ">=":
		if value > temperature.Min {
			temperature.Min = value
		}
	default:
		fmt.Printf("Invalid comparison sign '%s'\nThe temperature range has not changed\n", sign)
	}

	return temperature
}

func GetOptimalTemperature(temperature *TemperatureRange) int {
	if temperature.Min > temperature.Max {
		return -1
	}

	return temperature.Min
}
