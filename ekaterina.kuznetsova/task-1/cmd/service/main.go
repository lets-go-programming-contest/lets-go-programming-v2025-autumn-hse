package main

import (
	"fmt"
)

func main() {
	var (
		firstNum, secondNum int
		operation           string
	)
	_, err1 := fmt.Scanln(&firstNum)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err2 := fmt.Scanln(&secondNum)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, errOp := fmt.Scanln(&operation)
	if errOp != nil {
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
		fmt.Println(float64(firstNum) / float64(secondNum))
	default:
		fmt.Println("Invalid operation")
	}
}
