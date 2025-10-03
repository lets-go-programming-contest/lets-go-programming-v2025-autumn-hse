package main

import "fmt"

type Value struct {
	higher      int
	lower       int
	temperature int
}

const (
	MaxTemp = 30
	MinTemp = 15
)

func TValues() Value {
	return Value{
		higher:      MaxTemp,
		lower:       MinTemp,
		temperature: 0,
	}
}

func (values *Value) UpdateValues(operation string, temp int) {
	switch operation {
	case ">=":
		if temp > values.lower {
			values.lower = temp
		}

		if temp > values.temperature {
			values.temperature = temp
		}
	case "<=":
		if temp < values.higher {
			values.higher = temp
		}

		if temp < values.temperature {
			fmt.Print(-1)

			return
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
			fmt.Print("unable to read")

			return
		}

		values.UpdateValues(operation, temp)

		if values.temperature > values.higher {
			fmt.Println(-1)

			return
		}

		if values.temperature < values.lower {
			fmt.Println(-1)

			return
		}
	}

	fmt.Println(values.temperature)
}

func main() {
	var number, count int

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print("unable to read")

		return
	}

	for range number {
		_, err := fmt.Scan(&count)
		if err != nil {
			fmt.Print("unable to read")

			return
		}
		FindTemp(count)
	}
}
