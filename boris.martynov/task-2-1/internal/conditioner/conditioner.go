package conditioner

import (
	"fmt"
)

func TemperatureWantedDepartment() {
	var departmentCapacity, temperatureWantedByEmployee int64
	var greaterOrLess string
	lowestTemperature := int64(-30)
	highestTemperature := int64(99)
	fmt.Scanln(&departmentCapacity)
	for i := int64(0); i < departmentCapacity; i++ {
		fmt.Scanln(&greaterOrLess, &temperatureWantedByEmployee)
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

		if lowestTemperature < highestTemperature {
			fmt.Println(lowestTemperature)
		} else {
			fmt.Println("-1")
		}
	}
}
