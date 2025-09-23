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
    var op string
   
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

    fmt.Scan(&op)

    if op == "+" {
        s := a + b
        fmt.Println(s)
    } else if op == "-" {
        s := a - b
        fmt.Println(s)
    } else if op == "*" {
        s := a * b
        fmt.Println(s)
    } else if op == "/" {
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
