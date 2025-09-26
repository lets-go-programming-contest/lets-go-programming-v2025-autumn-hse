package conditioner

import (
	"fmt"
)

const (
	MinTemp = 0
	MaxTemp = 99
)

func TemperatureWantedDepartment() {
	var (
		departmentCapacity, temperatureWantedByEmployee int
		greaterOrLess                                   string
	)

	lowestTemperature := MinTemp
	lowestTemperatureSet := false
	highestTemperature := MaxTemp

	if _, err := fmt.Scanln(&departmentCapacity); err != nil {
		return
	}

	for range departmentCapacity {
		if _, err := fmt.Scanln(&greaterOrLess, &temperatureWantedByEmployee); err != nil {
			return
		}

		switch greaterOrLess {
		case ">=":
			if temperatureWantedByEmployee >= lowestTemperature {
				lowestTemperature = temperatureWantedByEmployee

			}

		case "<=":
			if temperatureWantedByEmployee <= highestTemperature {
				highestTemperature = temperatureWantedByEmployee
			}
		}

		switch lowestTemperatureSet {
		case true:
			if lowestTemperature < highestTemperature {
				fmt.Println(lowestTemperature)
			} else {
				fmt.Println(-1)
			}
		case false:
			fmt.Println(highestTemperature)
		}
	}
}
