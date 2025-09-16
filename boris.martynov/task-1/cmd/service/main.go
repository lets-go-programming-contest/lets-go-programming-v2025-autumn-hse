package main

import (
	"fmt"

	calc "github.com/JingolBong/task-1/internal/calculator"
)

func main() {
	firstNumber, secondNumber, operator := calc.UserInput()
	switch operator {
	case "+":
		fmt.Printf("%d\n", calc.Add(firstNumber, secondNumber))
	case "-":
		fmt.Printf("%d\n", calc.Subt(firstNumber, secondNumber))
	case "*":
		fmt.Printf("%d\n", calc.Mul(firstNumber, secondNumber))
	case "/":
		result, ok := calc.Div(firstNumber, secondNumber)
		if !ok {
			fmt.Println("Division by zero")
		} else {
			fmt.Printf("%d\n", result)
		}
	}
}
