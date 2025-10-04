package temperature

import "fmt"

type Value struct {
	higher int
	lower  int
}

const (
	MaxTemp = 30
	MinTemp = 15
)

func TValues() Value {
	return Value{
		higher: MaxTemp,
		lower:  MinTemp,
	}
}

func (values *Value) UpdateValues(operation string, temp int) {
	switch operation {
	case ">=":
		if temp > values.lower {
			values.lower = temp
		}
	case "<=":
		if temp < values.higher {
			values.higher = temp
		}
	default:
		fmt.Println("undefined operation")

		return
	}
}

func FindTemp(count int) {
	var (
		operation string
		temp      int
	)

	values := TValues()

	for range count {
		_, err := fmt.Scanln(&operation, &temp)
		if err != nil {
			fmt.Println("unable to read")

			return
		}

		values.UpdateValues(operation, temp)

		if values.lower <= values.higher {
			fmt.Println(values.lower)
		} else {
			fmt.Println(-1)
		}
	}
}
