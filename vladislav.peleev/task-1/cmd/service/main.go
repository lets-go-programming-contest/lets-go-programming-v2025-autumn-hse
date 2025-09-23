package main

import (
	"fmt"
	"strconv"
)

func readNumber() (int, error) {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return 0, fmt.Errorf("readNumber() : Scan error")
	}
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("readNumber() : Conversion error")
	}
	return num, nil
}

func main() {
	var a, b int
	var op string

	a, err := readNumber()
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	b, err = readNumber()
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&op)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch op {
	case "+":
		s := a + b
		fmt.Println(s)
	case "-":
		s := a - b
		fmt.Println(s)
	case "*":
		s := a * b
		fmt.Println(s)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		s := a / b
		fmt.Println(s)
	default:
		fmt.Println("Invalid operation")
		return
	}
}
