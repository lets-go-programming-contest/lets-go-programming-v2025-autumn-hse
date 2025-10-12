package temperature

import (
	"fmt"
	"io"
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

type preference struct {
	sign        string
	temperature int
}

func NewTemperatureProcessor() *TemperatureProcessor {
	return &TemperatureProcessor{
		lowerBound: minimumTemperature,
		upperBound: maximumTemperature,
	}
}

func readPreference(reader io.Reader) (preference, error) {
	var pref preference

	_, err := fmt.Fscan(reader, &pref.sign, &pref.temperature)
	if err != nil {
		return pref, err
	}

	return pref, nil
}

func printResult(result int) {
	fmt.Println(result)
}

func (tp *TemperatureProcessor) addPreference(pref preference) int {
	if pref.temperature < minimumTemperature || pref.temperature > maximumTemperature {
		return errorTemperature
	}

	switch pref.sign {
	case ">=":
		tp.lowerBound = max(tp.lowerBound, pref.temperature)
	case "<=":
		tp.upperBound = min(tp.upperBound, pref.temperature)
	}

	if tp.lowerBound > tp.upperBound {
		return errorTemperature
	}

	return tp.lowerBound
}

func (tp *TemperatureProcessor) ProcessDepartment(employeeCount int, reader io.Reader) error {
	for range employeeCount {
		pref, err := readPreference(reader)

		if err != nil {
			return fmt.Errorf("error reading sign and temperature: %w", err)
		}

		printResult(tp.addPreference(pref))
	}

	return nil
}
