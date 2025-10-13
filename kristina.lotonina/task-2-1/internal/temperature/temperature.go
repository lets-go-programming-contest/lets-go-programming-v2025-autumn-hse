package temperature

import "errors"

type Value struct {
	Higher int
	Lower  int
}

const (
	MaxTemp = 30
	MinTemp = 15
)

func NewValues() Value {
	return Value{
		Higher: MaxTemp,
		Lower:  MinTemp,
	}
}

func (values *Value) UpdateValues(operation string, temp int) error {
	switch operation {
	case ">=":
		if temp > values.Lower {
			values.Lower = temp
		}
	case "<=":
		if temp < values.Higher {
			values.Higher = temp
		}
	default:
		return errors.New("undefined operation :" + operation)
	}
	return nil
}
