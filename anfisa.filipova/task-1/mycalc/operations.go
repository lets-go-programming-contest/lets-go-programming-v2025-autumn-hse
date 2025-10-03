package mycalc

import (
	"errors"
)

var (
	errDivisionByZero   = errors.New("Division by zero")
	errInvalidOperation = errors.New("Invalid operation")
)

func Calculate(operand1 int, operand2 int, operation string) (int, error) {
	switch operation {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, errDivisionByZero
		}
		return operand1 / operand2, nil
	default:
		return 0, errInvalidOperation
	}
}
