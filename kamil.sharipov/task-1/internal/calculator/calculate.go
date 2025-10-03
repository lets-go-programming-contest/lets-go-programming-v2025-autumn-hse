package calculator

import (
	"errors"
)

var (
	ErrDivisionByZero   = errors.New("Division by zero")
	ErrInvalidOperation = errors.New("Invalid operation")
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
			return 0, ErrDivisionByZero
		}
		return int64(leftOperand / rightOperand), nil
	default:
		return 0, ErrInvalidOperation
	}
}
