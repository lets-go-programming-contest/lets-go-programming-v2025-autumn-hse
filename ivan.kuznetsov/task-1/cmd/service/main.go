package main

import "fmt"

func main() {
	var num1, num2 int
	var operation string

	_, err1 := fmt.Scan(&num1)
	_, err2 := fmt.Scan(&num2)
	_, err3 := fmt.Scan(&operation)

	if err1 != nil {
		fmt.Println("Invalid first operand")
	} else if err2 != nil {
		fmt.Println("Invalid second operand")
	} else if err3 != nil {
		fmt.Println("Invalid operation")
	} else {
		switch operation {
		case "+":
			fmt.Println(num1 + num2)
		case "-":
			fmt.Println(num1 - num2)
		case "*":
			fmt.Println(num1 * num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Division by zero")
			} else {
				fmt.Println(num1 / num2)
			}
		default:
			fmt.Println("Invalid operation")
		}
	}
}
