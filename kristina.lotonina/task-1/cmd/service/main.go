package main

import (
	"fmt"

	"github.com/kef1rch1k/task-1/internal/calculator"
)

func main() {
	var (
		first, second int
		operation     string
	)
	_, err := fmt.Scan(&first)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&second)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}
	calculator.Calculate(first, second, operation)
}
