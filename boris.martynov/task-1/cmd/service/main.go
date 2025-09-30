package main

import (
	"fmt"

	"errors"

	"github.com/JingolBong/task-1/internal/inputreader"
)

var errDivisionByZero = errors.New("Division by zero")

func main() {
	firstNumber, err := inputreader.ReadNumber("first operand")
	if err != nil {
		fmt.Println(err)
		return
	}
	secondNumber, err := inputreader.ReadNumber("second operand")
	if err != nil {
		fmt.Println(err)
		return
	}
	operator, err := inputreader.ReadOperator()
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
		if secondNumber == 0 {
			fmt.Println(errDivisionByZero)
			return
		}
		fmt.Println(firstNumber / secondNumber)
	default:
		fmt.Println(inputreader.ErrInvalidOperation)
	}
}
