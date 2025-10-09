package main

import (
	"fmt"

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
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}
