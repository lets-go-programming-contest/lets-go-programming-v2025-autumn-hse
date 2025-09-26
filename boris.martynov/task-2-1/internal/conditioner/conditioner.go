package conditioner

import (
	"fmt"
)

const (
	MinTemp = 15
	MaxTemp = 30
)

func TemperatureWantedDepartment() {
	var (
		departmentCapacity, temperatureWantedByEmployee int
		greaterOrLess                                   string
	)

	lowestTemperature := MinTemp
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

		if lowestTemperature <= highestTemperature {
			fmt.Println(lowestTemperature)
		} else {
			fmt.Println("-1")
		}
	}
}
