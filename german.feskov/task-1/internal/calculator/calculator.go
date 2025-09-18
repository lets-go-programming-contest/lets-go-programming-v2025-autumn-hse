package calculator

import (
	"fmt"
)

type Operand int

type Operation string

func Calculate(a, b Operand, op Operation) (Operand, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Invalid operation")
	}
}
