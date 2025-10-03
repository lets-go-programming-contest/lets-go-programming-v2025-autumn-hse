package main

import "fmt"

type Value struct {
	higher      int
	lower       int
	temperature int
}

func TValues() Value {
	return Value{
		higher:      30,
		lower:       15,
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
	var operation string
	var temp 	  int
	values := TValues()
	for range count {
		fmt.Scanln(&operation, &temp)
		values.UpdateValues(operation, temp)
		if values.temperature > values.higher {
			fmt.Print(-1)

			return
		}
		if values.temperature < values.lower {
			fmt.Print(-1)

			return
		}
	}
	fmt.Println(values.temperature)	
}

func main() {
	var (
		number, count int
	)
	fmt.Scan(&number, &count)
	for range number {
		FindTemp(count)
	}
}
