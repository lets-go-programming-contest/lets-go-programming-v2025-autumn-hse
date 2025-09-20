package main

import (
	"fmt"
	"strings"

	"github.com/kamilSharipov/task-1/internal/calculator"
)

func main() {
	var (
		leftOperand, rightOperand int32
		operator                  string
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

	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	result, err := calculator.Calculate(leftOperand, rightOperand, operator)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "division by zero"):
			fmt.Println("Division by zero")
		case strings.Contains(err.Error(), "invalid operation"):
			fmt.Println("Invalid operation")
		default:
			return
		}
	} else {
		fmt.Println(result)
	}
}
