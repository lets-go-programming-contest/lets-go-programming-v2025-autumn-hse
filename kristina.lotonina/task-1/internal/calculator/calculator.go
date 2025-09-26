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
			fmt.Print("Division by zero\n")	
		}
	} else {
		fmt.Print("Invalid operation\n")
	}
}
