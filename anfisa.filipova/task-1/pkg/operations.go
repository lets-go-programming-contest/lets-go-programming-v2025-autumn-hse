package pkg

import "fmt"

func Calculate(operand1 int, operand2 int, operation string) (interface{}, error) {
	switch operation {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return float64(operand1) / float64(operand2), nil
	default:
		return 0, fmt.Errorf("Invalid operation")
	}

}
