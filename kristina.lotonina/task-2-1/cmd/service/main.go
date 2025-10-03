package main

import "fmt"

type Value struct {
	higher string
	lower  string
}

func TValues() Value {
	return Value{
		higher: "30",
		lower:  "15",
	}
}

func main() {
	var (
		number, count int
		input, t      string
	)

	t = "0"
	values := TValues()
	fmt.Scan(&number, &count)

	for i := range number {
		for j := range count {
			fmt.Scanln(&input)
			if string(input[0])+string(input[1]) == ">=" {
				if string(input[2])+string(input[3]) > values.lower {
					values.lower = string(input[2]) + string(input[3])
				}
				if t < string(input[2])+string(input[3]) {
					t = string(input[2]) + string(input[3])
					if t > values.higher {
						fmt.Print(-1)
						return
					}
					if t < values.lower {
						fmt.Print(-1)
						return
					}
				}
			}
			if string(input[0])+string(input[1]) == "<=" {
				if string(input[2])+string(input[3]) < values.higher {
					values.higher = string(input[2]) + string(input[3])
				}
				if string(input[2])+string(input[3]) < t {
					fmt.Print(-1)
					return
				}
				if t > values.higher {
					fmt.Print(-1)
					return
				}
				if t < values.lower {
					fmt.Print(-1)
					return
				}
			}
			fmt.Println(t)
		}
	}
}
