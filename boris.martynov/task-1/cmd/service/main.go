package main

import (
	"fmt"
	"github.com/JingolBong/task-1/internal/calculator"
)

func main() {
	firstNumber, secondNumber, operator := calc.UserInput()
	switch operator {
	case "+":
		fmt.Printf("Result: %d\n", calc.Add(firstNumber, secondNumber))
	case "-":
		fmt.Printf("Result: %d\n", calc.Subt(firstNumber, secondNumber))
	case "*":
		fmt.Printf("Result: %d\n", calc.Mul(firstNumber, secondNumber))
	case "/":
		fmt.Printf("Result: %.2f\n", calc.Div(firstNumber, secondNumber))
	}
}
