package mycalc

import (
	"errors"
)

var errDivisionByZero = errors.New("Division by zero")
var errInvalidOperation = errors.New("Invalid operation")

func Calculate(operand1 int, operand2 int, operation string) (float64, error) {
	switch operation {
	case "+":
		return float64(operand1 + operand2), nil
	case "-":
		return float64(operand1 - operand2), nil
	case "*":
		return float64(operand1 * operand2), nil
	case "/":
		if operand2 == 0 {
			return 0, errDivisionByZero
		}
		return float64(operand1) / float64(operand2), nil
	default:
		return 0, errInvalidOperation
	}
}
