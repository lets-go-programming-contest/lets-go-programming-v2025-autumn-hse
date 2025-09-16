package calculator

import (
	"fmt"
)

func Calculate(left_operand, right_operand int32, operator rune) int64 {
	switch operator {
	case '+':
		return int64(left_operand + right_operand)
	case '-':
		return int64(left_operand - right_operand)
	case '*':
		return int64(left_operand * right_operand)
	case '/':
		if right_operand == 0 {
			fmt.Println("Division by zero")
			return 0
		}
		return int64(left_operand / right_operand)
	default:
		fmt.Println("Invalid operation")
		return 0
	}

	return 0
}
