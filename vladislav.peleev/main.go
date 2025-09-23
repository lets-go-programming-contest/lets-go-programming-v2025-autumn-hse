package main

import (
    "fmt"
    "strconv"
)

func read_number() (int, error) {
    var input string
    fmt.Scan(&input)
    num, err := strconv.Atoi(input)
    if err != nil {
        return 0, fmt.Errorf("")
    }
    return num, nil
}

func main() {
    var a, b int
    var operand string
   
    a, err := read_number()
    if err != nil {
        fmt.Println("Invalid first operand")
        return
    }
    
    b, err = read_number()
    if err != nil {
        fmt.Println("Invalid second operand")
	return
    }

    fmt.Scan(&operand)

    if operand == "+" {
        s := a + b
        fmt.Println(s)
    } else if operand == "-" {
        s := a - b
        fmt.Println(s)
    } else if operand == "*" {
        s := a * b
        fmt.Println(s)
    } else if operand == "/" {
        if b == 0 {
            fmt.Println("Division by zero")
            return
        } else {
            s := a / b
            fmt.Println(s)
        }
    } else {
        fmt.Println("Invalid operation")
        return
    }
}
