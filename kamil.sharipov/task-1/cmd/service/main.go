package main

import (
	"fmt"

	"github.com/kamilSharipov/task-1/internal/calculator"
)

func main() {
	var leftOperand, rightOperand int32
	var operator string

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

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Invalid operation")
		return
	}

	result, ok := calculator.Calculate(
		leftOperand, rightOperand, rune(operator[0]))
	if ok {
		fmt.Printf("%d\n", result)
	}
}
