package temperature

import (
	"errors"
	"fmt"
	"io"

	"github.com/Tapochek2894/task-2/subtask-1/internal/intmath"
)

var (
	errDivision         = errors.New("Division by zero")
	errInvalidOperation = errors.New("Invalid operation")
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

func (tp *TemperatureProcessor) ProcessDepartment(reader io.Reader) ([]int, error) {
	var employeeCount int

	_, err := fmt.Fscan(reader, &employeeCount)
	if err != nil {
		return nil, fmt.Errorf("error reading employee count: %w", err)
	}

	resultSlice := make([]int, employeeCount)

	for i := range employeeCount {
		var (
			sign        string
			temperature int
		)

		_, err := fmt.Fscan(reader, &sign, &temperature)
		if err != nil {
			return nil, fmt.Errorf("error reading sign and temperature: %w", err)
		}

		resultSlice[i] = tp.addPreference(sign, temperature)
	}

	return resultSlice, nil
}
