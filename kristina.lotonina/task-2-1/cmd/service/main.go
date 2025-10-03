package main

import "fmt"

type Value struct {
	higher      string
	lower       string
	temperature string
}

func TValues() Value {
	return Value{
		higher:      "30",
		lower:       "15",
		temperature: "0",
	}
}

func FindTemp(number int, count int) {
	values := TValues()
	var input string
	for range number {
		for range count {
			fmt.Scanln(&input)
			operator := string(input[0]) + string(input[1])
			switch operator {
			case ">=":
				if string(input[2])+string(input[3]) > values.lower {
					values.lower = string(input[2]) + string(input[3])
				}
				if values.temperature < string(input[2])+string(input[3]) {
					values.temperature = string(input[2]) + string(input[3])
					if values.temperature > values.higher {
						fmt.Print(-1)
						return
					}
					if values.temperature < values.lower {
						fmt.Print(-1)
						return
					}
				}
			case "<=":
				if string(input[2])+string(input[3]) < values.higher {
					values.higher = string(input[2]) + string(input[3])
				}
				if string(input[2])+string(input[3]) < values.temperature {
					fmt.Print(-1)
					return
				}
				if values.temperature > values.higher {
					fmt.Print(-1)
					return
				}
				if values.temperature < values.lower {
					fmt.Print(-1)
					return
				}
			default:
				fmt.Println("undefined operation")
				return
			}
		}
	}
	fmt.Println(values.temperature)
}

func main() {
	var (
		number, count int
	)
	fmt.Scan(&number, &count)
	FindTemp(number, count)
}
