package calculator

import (
	"fmt"
)

func Calculate(leftOperand, rightOperand int32, operator string) (int64, error) {
	switch operator {
	case "+":
		return int64(leftOperand) + int64(rightOperand), nil
	case "-":
		return int64(leftOperand) - int64(rightOperand), nil
	case "*":
		return int64(leftOperand) * int64(rightOperand), nil
	case "/":
		if rightOperand == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return int64(leftOperand / rightOperand), nil
	default:
		return 0, fmt.Errorf("invalid operation")
	}
}
