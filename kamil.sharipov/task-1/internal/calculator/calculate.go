package calculator

import (
	"fmt"
)

func Calculate(left_operand, right_operand int32, operator rune) (int64, bool) {
	switch operator {
	case '+':
		return int64(left_operand) + int64(right_operand), true
	case '-':
		return int64(left_operand) - int64(right_operand), true
	case '*':
		return int64(left_operand) * int64(right_operand), true
	case '/':
		if right_operand == 0 {
			fmt.Println("Division by zero")
			return 0, false
		}
		return int64(left_operand / right_operand), true
	default:
		fmt.Println("Invalid operation")
		return 0, false
	}
}
