package main

import (
	"fmt"
	"strings"
)

func main() {
	var first_operand, second_operand int64
	var result int64
	var operation string

	_, err_first_operand := fmt.Scan(&first_operand)
	_, err_second_operand := fmt.Scan(&second_operand)
	_, err_operation := fmt.Scan(&operation)

	if err_first_operand != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if err_second_operand != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if err_operation != nil {
		fmt.Println("Invalid operation")
		return
	}
	if !strings.Contains("+-*/", operation) {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		result = first_operand + second_operand
	case "-":
		result = first_operand - second_operand
	case "*":
		result = first_operand * second_operand
	case "/":
		if second_operand == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = first_operand / second_operand
	}

	fmt.Println(result)
}
