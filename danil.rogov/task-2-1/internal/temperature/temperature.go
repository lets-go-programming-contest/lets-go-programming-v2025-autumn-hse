package temperature

import (
	"errors"
	"fmt"
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

type Preference struct {
	Sign        string
	Temperature int
}

var ErrInvalidSign = errors.New("invalid sign")

func NewTemperatureProcessor() *TemperatureProcessor {
	return &TemperatureProcessor{
		lowerBound: minimumTemperature,
		upperBound: maximumTemperature,
	}
}

func (tp *TemperatureProcessor) AddPreference(pref Preference) (int, error) {
	if pref.Temperature < minimumTemperature || pref.Temperature > maximumTemperature {
		return errorTemperature, nil
	}

	switch pref.Sign {
	case ">=":
		tp.lowerBound = max(tp.lowerBound, pref.Temperature)
	case "<=":
		tp.upperBound = min(tp.upperBound, pref.Temperature)
	default:
		return 0, fmt.Errorf("%w: %s", ErrInvalidSign, pref.Sign)
	}

	if tp.lowerBound > tp.upperBound {
		return errorTemperature, nil
	}

	return tp.lowerBound, nil
}
