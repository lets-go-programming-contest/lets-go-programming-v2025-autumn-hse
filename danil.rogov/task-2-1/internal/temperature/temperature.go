package temperature

import (
	"fmt"
	"io"

	"github.com/Tapochek2894/task-2/subtask-1/internal/intmath"
)

const (
	minimumTemperature = 15
	maximumTemperature = 30
	errorTemperature   = -1
)

type TemperatureProcessor struct {
	lowerBound int
	upperBound int
}

func NewTemperatureProcessor() *TemperatureProcessor {
	return &TemperatureProcessor{
		lowerBound: minimumTemperature,
		upperBound: maximumTemperature,
	}
}

func (tp *TemperatureProcessor) addPreference(sign string, temperature int) int {
	if temperature < minimumTemperature || temperature > maximumTemperature {
		return errorTemperature
	}

	switch sign {
	case ">=":
		tp.lowerBound = intmath.LargerInt(tp.lowerBound, temperature)
	case "<=":
		tp.upperBound = intmath.SmallerInt(tp.upperBound, temperature)
	}

	if tp.lowerBound > tp.upperBound {
		return errorTemperature
	}

	return tp.lowerBound
}

func (tp *TemperatureProcessor) ProcessDepartment(reader io.Reader) {
	var employeeCount int

	_, err := fmt.Fscan(reader, &employeeCount)
	if err != nil {
		fmt.Println("Error reading employee count:", err)

		return
	}

	for range employeeCount {
		var (
			sign        string
			temperature int
		)

		_, err := fmt.Fscan(reader, &sign, &temperature)
		if err != nil {
			fmt.Println("Error reading sign and temperature:", err)

			return
		}

		result := tp.addPreference(sign, temperature)
		fmt.Println(result)
	}
}
