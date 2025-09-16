package calculator

import "errors"

func Calculate(firstOperand, secondOperand int64, operation string) (int64, error) {
	switch operation {
	case "+":
		return firstOperand + secondOperand, nil
	case "-":
		return firstOperand - secondOperand, nil
	case "*":
		return firstOperand * secondOperand, nil
	case "/":
		if secondOperand == 0 {
			return 0, errors.New("Division by zero")
		} else {
			return firstOperand / secondOperand, nil
		}
	default:
		return 0, errors.New("Invalid operation")
	}
}
