package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b float64
	var op string

	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid first operand")
		return
	}
	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid second operand")
		return
	}
	if _, err := fmt.Scanln(&op); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid operation")
		return
	}

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Fprintln(os.Stderr, "Division by zero")
			return
		}
		fmt.Println(a / b)
	default:
		fmt.Fprintln(os.Stderr, "Invalid operation")
	}
}
