package conditioner

import (
	"fmt"
)

const (
	MinTemp = -30
	MaxTemp = 99
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
	for i := 0; i < departmentCapacity; i++ {
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

		if lowestTemperature > highestTemperature {
			fmt.Println("-1")
		} else {
			fmt.Println(lowestTemperature)
		}

	}
}
