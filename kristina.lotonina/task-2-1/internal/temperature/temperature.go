package temperature

import (
	"errors"
	"fmt"
)

var ErrUndefinedOperation = errors.New("undefined operation")

type Value struct {
	Higher int
	Lower  int
}

<<<<<<< HEAD
func NewValues(higher, lower int) Value {
	return Value{
		Higher: higher,
		Lower:  lower,
	}
}

func (v* Value) FindTemp(operation string, temp int) (int, error) {
	err := v.UpdateValues(operation, temp)
	if err != nil {
		return 0, fmt.Errorf("failed to update temperature values: %w", err)
	}

	if v.Lower <= v.Higher {
		return v.Lower, nil
	}

	return -1, nil
}

=======
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

>>>>>>> b738d5c9f7fb824b1236f4c6877627be159127ef
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
		return fmt.Errorf("%w: %s", ErrUndefinedOperation, operation)
	}

	return nil
}
