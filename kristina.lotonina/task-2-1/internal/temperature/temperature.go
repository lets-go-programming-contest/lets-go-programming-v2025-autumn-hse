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

func NewValues(higher, lower int) Value {
	return Value{
		Higher: higher,
		Lower:  lower,
	}
}

func (v *Value) UpdateTemp(operation string, temp int) error {
    return v.UpdateValues(operation, temp)
}

func (v *Value) GetCurrentTemp() int {
    if v.Lower <= v.Higher {
        return v.Lower
    }
    return -1
}

func (v *Value) UpdateValues(operation string, temp int) error {
	switch operation {
	case ">=":
		if temp > v.Lower {
			v.Lower = temp
		}
	case "<=":
		if temp < v.Higher {
			v.Higher = temp
		}
	default:
		return fmt.Errorf("%w: %s", ErrUndefinedOperation, operation)
	}

	return nil
}
