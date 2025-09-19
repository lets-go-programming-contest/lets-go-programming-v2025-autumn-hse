package calculator

import (
	"fmt"
)

func Calculate(leftOperand, rightOperand int32, operator string) (int64, bool) {
	switch operator {
	case "+":
		return int64(leftOperand) + int64(rightOperand), true
	case "-":
		return int64(leftOperand) - int64(rightOperand), true
	case "*":
		return int64(leftOperand) * int64(rightOperand), true
	case "/":
		return int64(leftOperand / rightOperand), true
	default:
		fmt.Println("Invalid operation")
		return 0, false
	}
}
