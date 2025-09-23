package main

import (
	"fmt"
)

func main() {
	var (
		firstNum, secondNum int
		operation           string
	)
	_, err := fmt.Scanln(&firstNum)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&secondNum)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(firstNum + secondNum)
	case "-":
		fmt.Println(firstNum - secondNum)
	case "*":
		fmt.Println(firstNum * secondNum)
	case "/":
		if secondNum == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(firstNum / secondNum)
	default:
		fmt.Println("Invalid operation")
	}
}
