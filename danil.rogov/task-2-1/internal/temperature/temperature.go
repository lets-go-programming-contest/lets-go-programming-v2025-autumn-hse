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
		return preference{}, fmt.Errorf("failed to read preference: %w", err)
	}

	return pref, nil
}

func printResult(result int, writer io.Writer) error {
	_, err := fmt.Fprintln(writer, result)
	if err != nil {
		return fmt.Errorf("failed to write result: %w", err)
	}

	return nil
}

func (tp *TemperatureProcessor) addPreference(pref preference) (int, error) {
	if pref.temperature < minimumTemperature || pref.temperature > maximumTemperature {
		return errorTemperature, nil
	}

	switch pref.sign {
	case ">=":
		tp.lowerBound = max(tp.lowerBound, pref.temperature)
	case "<=":
		tp.upperBound = min(tp.upperBound, pref.temperature)
	default:
		return 0, fmt.Errorf("invalid sign: %s", pref.sign)
	}

	if tp.lowerBound > tp.upperBound {
		return errorTemperature, nil
	}

	return tp.lowerBound, nil
}

func (tp *TemperatureProcessor) ProcessDepartment(
	employeeCount int,
	reader io.Reader,
	writer io.Writer,
) error {
	for range employeeCount {
		preference, err := readPreference(reader)
		if err != nil {
			return err
		}

		newPreference, err := tp.addPreference(preference)
		if err != nil {
			return err
		}

		err = printResult(newPreference, writer)
		if err != nil {
			return err
		}
	}

	return nil
}
