package main

import (
	"fmt"

	"errors"

	"github.com/JingolBong/task-1/internal/inputreader"
)

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
			fmt.Println(errors.New("Division by zero"))
			return
		}
		fmt.Println(firstNumber / secondNumber)
	default:
		fmt.Println(errors.New("Invalid operation"))
	}
}
