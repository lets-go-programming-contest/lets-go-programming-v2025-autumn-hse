package main

import (
	"fmt"

	"github.com/Anfisa111/task-1/pkg"
)

func main() {
	var operation string
	var num1, num2 int
	_, err := fmt.Scan(&num1)

	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&num2)

	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, _ = fmt.Scan(&operation)

	result, err := pkg.Calculate(num1, num2, operation)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
