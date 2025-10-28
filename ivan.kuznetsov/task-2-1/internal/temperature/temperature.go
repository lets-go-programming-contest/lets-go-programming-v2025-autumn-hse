package temperature

import "errors"

var ErrInvalidSign = errors.New("invalid comparison sign")

const (
	MinTemperature = 15
	MaxTemperature = 30
)

type TemperatureRange struct {
	Min int
	Max int
}

func NewTemperatureRange(minTvalue int, maxTvalue int) *TemperatureRange {
	return &TemperatureRange{
		Min: minTvalue,
		Max: maxTvalue,
	}
}

func (temperature *TemperatureRange) OptimalTemperature(sign string, value int) error {
	switch sign {
	case "<=":
		if value < temperature.Max {
			temperature.Max = value
		}
	case ">=":
		if value > temperature.Min {
			temperature.Min = value
		}
	default:
		return ErrInvalidSign
	}

	return nil
}

func (temperature *TemperatureRange) GetOptimalTemperature() int {
	if temperature.Min > temperature.Max {
		return -1
	}

	return temperature.Min
}
