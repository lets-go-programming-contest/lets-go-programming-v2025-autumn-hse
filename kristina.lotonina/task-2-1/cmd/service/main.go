package main

import "fmt"

type Value struct {
	higher      int
	lower       int
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
		}
	}
}

func main() {
	var number, count int

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("unable to read")

		return
	}

	for range number {
		_, err = fmt.Scan(&count)
		if err != nil {
			fmt.Println("unable to read")

			return
		}

		FindTemp(count)
	}
}
