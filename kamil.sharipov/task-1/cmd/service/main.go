package main

import (
	"fmt"

	"github.com/kamilSharipov/task-1/internal/calculator"
)

func main() {
	var (
		leftOperand  int32
		rightOperand int32
		operator     string
	)

	_, err := fmt.Scanln(&leftOperand)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&rightOperand)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if rightOperand == 0 {
		fmt.Println("Division by zero")
	}

	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	result, ok := calculator.Calculate(
		leftOperand, rightOperand, operator)
	if ok {
		fmt.Printf("%d\n", result)
	}
}
