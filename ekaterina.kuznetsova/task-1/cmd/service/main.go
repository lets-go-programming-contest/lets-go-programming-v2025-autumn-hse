package main 

import "fmt"

func main() {
	var first_num, second_num int
	var operation string
	_, err_1 := fmt.Scanln(&first_num)
	if err_1 != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err_2 := fmt.Scanln(&second_num)
	if err_2 != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err_op := fmt.Scanln(&operation)
	if err_op != nil {
		fmt.Println("Invalid operation")
		return
	}

	if operation == "/" && second_num == 0 {
		fmt.Println("Division by zero")
		return
	}

	switch operation {
	case "+":
		fmt.Println(first_num + second_num)
	case "-":
		fmt.Println(first_num - second_num)
	case "*":
		fmt.Println(first_num * second_num)
	case "/":
		fmt.Println(float64(first_num) / float64(second_num))
	default:
		fmt.Println("Invalid operation")
	}
}
