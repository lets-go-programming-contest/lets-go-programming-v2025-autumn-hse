package main

import (
	"fmt"

	"github.com/JingolBong/task-1/internal/calculator"
)

func main() {
	firstNumber, err := calculator.ReadNumber("first operand")
	if err != nil {
		fmt.Println(err)
		return
	}
	secondNumber, err := calculator.ReadNumber("second operand")
	if err != nil {
		fmt.Println(err)
		return
	}
	operator, err := calculator.ReadOperator()
	if err != nil {
		fmt.Println(err)
		return
	}
	switch operator {
	case "+":
		fmt.Println(firstNumber + secondNumber)
	case "-":
		fmt.Println(firstNumber - secondNumber)
	case "*":
		fmt.Println(firstNumber * secondNumber)
	case "/":
		result, err := calculator.Div(firstNumber, secondNumber)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
