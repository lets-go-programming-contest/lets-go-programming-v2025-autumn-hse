package calculator

import "fmt"

func calculate(a int, b int, op string) {
	if (op == "+") {
		fmt.Print(a + b) 
	} else if (op == "-") {
		fmt.Print(a - b)
	} else if (op == "*") {
		fmt.Print(a * b)
	} else if (op == "/") {
		if (b != 0) {
			fmt.Print(a / b)
		} else {
			fmt.Print("Division by zero\n")	
		}
	} else {
		fmt.Print("Invalid operation\n")
	}
}
