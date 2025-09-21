package temperature

import (
	"fmt"

	"github.com/Tapochek2894/task-2/subtask-1/internal/intmath"
)

const (
	minimumTemperature = 15
	maximumTemperature = 30
	errorTemperature   = -1
)

func SetDepartmentTemperature() {
	var employeeCount int
	_, err := fmt.Scan(&employeeCount)

	if err != nil {
		fmt.Println("Error reading employee count:", err)

		return
	}

	lowerBound := minimumTemperature
	upperBound := maximumTemperature

	for range employeeCount {
		var (
			preferredTemperature int
			inequalitySign       string
		)

		_, err = fmt.Scan(&inequalitySign, &preferredTemperature)
		if err != nil {
			fmt.Println("Error reading sign or temperature:", err)

			return
		}

		if preferredTemperature < minimumTemperature || preferredTemperature > maximumTemperature {
			fmt.Println(errorTemperature)

			continue
		}

		switch inequalitySign {
		case ">=":
			lowerBound = intmath.LargerInt(lowerBound, preferredTemperature)
		case "<=":
			upperBound = intmath.SmallerInt(upperBound, preferredTemperature)
		}

		result := lowerBound
		if lowerBound > upperBound {
			result = errorTemperature
		}

		fmt.Println(result)
	}
}
