package calculator

import "fmt"

func Calculate(a int, b int, op string) {
	if op == "+" {
		fmt.Println(a + b) 
	} else if op == "-" {
		fmt.Println(a - b)
	} else if op == "*" {
		fmt.Println(a * b)
	} else if op == "/" {
		if b != 0 {
			fmt.Println(a / b)
		} else {
			fmt.Println("Division by zero")	
		}
	} else {
		fmt.Println("Invalid operation")
	}
}
