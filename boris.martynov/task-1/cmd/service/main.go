package main

import (
	"fmt"

	"errors"

	"github.com/JingolBong/task-1/internal/input_reader"
)

func main() {
	firstNumber, err := input_reader.ReadNumber("first operand")
	if err != nil {
		fmt.Println(err)
		return
	}
	secondNumber, err := input_reader.ReadNumber("second operand")
	if err != nil {
		fmt.Println(err)
		return
	}
	operator, err := input_reader.ReadOperator()
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
			fmt.Println(errors.New("division by zero"))
			return
		}
		fmt.Println(firstNumber / secondNumber)
	default:
		fmt.Println(errors.New("Invalid operation"))
	}
}
