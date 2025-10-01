package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    var first_operand int
    _, err_first := fmt.Scan(&first_operand)
    if err_first != nil {
        fmt.Println("Invalid first operand")
        bufio.NewReader(os.Stdin).ReadString('\n')
        return
    }
    
    var second_operand int
    _, err_second := fmt.Scan(&second_operand)
    if err_second != nil {
        fmt.Println("Invalid second operand")
        bufio.NewReader(os.Stdin).ReadString('\n')
        return
    }
    
    var operator string
    _, _ = fmt.Scan(&operator)
    /*if err_op != nil {
        fmt.Println("Invalid operator")
    }*/
    
    switch operator {
    
    case "+":
        fmt.Println(first_operand + second_operand)
    case "-":
        fmt.Println(first_operand - second_operand)
    case "*":
        fmt.Println(first_operand * second_operand)
    case "/":
        if second_operand == 0 {
            fmt.Println("Division by zero")
        } else {
            fmt.Println(first_operand / second_operand)
        }
    default:
        fmt.Println("Invalid operator")
        bufio.NewReader(os.Stdin).ReadString('\n')
        return
    }
}
