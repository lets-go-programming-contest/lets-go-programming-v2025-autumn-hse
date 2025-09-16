package main

import (
	"fmt"

	"github.com/kamilSharipov/task-1/internal/calculator"
)

func main() {
	var left_operand, right_operand int32
	var operator string

	_, err := fmt.Scanln(&left_operand)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&right_operand)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Invalid operation")
		return
	}

	result := calculator.Calculate(left_operand, right_operand, rune(operator[0]))
	fmt.Printf("%d\n", result)
}
