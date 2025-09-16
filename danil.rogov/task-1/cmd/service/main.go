package main

import (
	"fmt"

	"github.com/Tapochek2894/task-1/internal/calculator"
)

func main() {
	var (
		firstOperand, secondOperand int64
		operation                   string
	)
	_, err := fmt.Scan(&firstOperand)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&secondOperand)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}
	result, err := calculator.Calculate(firstOperand, secondOperand, operation)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
