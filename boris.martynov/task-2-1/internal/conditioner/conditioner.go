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
			if !lowestTemperatureSet || temperatureWantedByEmployee >= lowestTemperature {
				lowestTemperature = temperatureWantedByEmployee
				lowestTemperatureSet = true
			}

		case "<=":
			if temperatureWantedByEmployee <= highestTemperature {
				highestTemperature = temperatureWantedByEmployee
			}
		}

		if lowestTemperatureSet && (lowestTemperature < highestTemperature) {
			fmt.Println(lowestTemperature)
		} else {
			fmt.Println("-1")
		}
	}
}
