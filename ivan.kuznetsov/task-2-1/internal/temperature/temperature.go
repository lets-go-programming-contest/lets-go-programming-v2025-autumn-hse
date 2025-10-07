package temperature

type TemperatureRange struct {
	Min int
	Max int
}

const (
	MinTemperature = 15
	MaxTemperature = 30
)

func TemperatureRangeInit() *TemperatureRange {
	return &TemperatureRange{Min: MinTemperature, Max: MaxTemperature}
}

func OptimalTemperature(sign string, value int, temperature *TemperatureRange) *TemperatureRange {
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
		return nil
	}

	return temperature
}

func GetOptimalTemperature(temperature *TemperatureRange) int {
	if temperature.Min > temperature.Max {
		return -1
	}

	return temperature.Min
}
