package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstOperand, secondOperand int64
	var result int64
	var operation string

	_, errFirstOperand := fmt.Scan(&firstOperand)
	_, errSecondOperand := fmt.Scan(&secondOperand)
	_, errOperation := fmt.Scan(&operation)

	if errFirstOperand != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if errSecondOperand != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if errOperation != nil {
		fmt.Println("Invalid operation")
		return
	}
	if !strings.Contains("+-*/", operation) {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		result = firstOperand + secondOperand
	case "-":
		result = firstOperand - secondOperand
	case "*":
		result = firstOperand * secondOperand
	case "/":
		if secondOperand == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = firstOperand / secondOperand
	}

	fmt.Println(result)
}
