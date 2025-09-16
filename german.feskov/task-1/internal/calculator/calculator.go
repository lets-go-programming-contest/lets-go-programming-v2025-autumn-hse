package calculator

import (
	"fmt"
)

type Operand int

type Operation string

func Calculate(a, b Operand, op Operation) (Operand, error) {
	var (
		result Operand
		err    error
	)
	switch op {
	case "+":
		result, err = a+b, nil
	case "-":
		result, err = a-b, nil
	case "*":
		result, err = a*b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		result, err = a/b, nil
	default:
		return 0, fmt.Errorf("Invalid operation")
	}
	return result, err
}
