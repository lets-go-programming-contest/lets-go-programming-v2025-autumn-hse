package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		firstOperand, secondOperand float64
		sign                        string
	)
	_, err := fmt.Scan(&firstOperand)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid first operand")
		os.Exit(1)
	}
	_, err = fmt.Scan(&secondOperand)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid second operand")
		os.Exit(1)
	}
	_, err = fmt.Scan(&sign)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid operand")
		os.Exit(1)
	}
	switch sign {
	case "+":
		fmt.Println(firstOperand + secondOperand)
	case "-":
		fmt.Println(firstOperand - secondOperand)
	case "*":
		fmt.Println(firstOperand * secondOperand)
	case "/":
		if secondOperand == 0 {
			fmt.Fprintln(os.Stderr, "Division by zero")
			os.Exit(1)
		}
		fmt.Println(firstOperand / secondOperand)
	default:
		fmt.Fprintln(os.Stderr, "Invalid operand")
		os.Exit(1)
	}
}
