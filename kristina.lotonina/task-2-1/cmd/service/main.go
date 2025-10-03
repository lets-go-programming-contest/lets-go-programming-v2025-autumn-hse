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
		n, k     int
		input, t string
	)

	t = "0"
	values := TValues()
	fmt.Scan(&n, &k)

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
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
