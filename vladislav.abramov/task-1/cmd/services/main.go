package main

import (
  "fmt"
)

func main() {
  var a int
  var b int
  var op string

  _, error1 := fmt.Scanln(&a)
  if error1 != nil {
    fmt.Print("Invalid first operand\n")
    return
  }
  _, error2 := fmt.Scanln(&b)
  if error2 != nil {
    fmt.Print("Invalid second operand\n")
    return
  }
  _, error3 := fmt.Scanln(&op)
  if error3 != nil {
    fmt.Print("Invalid operation\n")
    return
  }
  if b == 0 && op == "/" {
    fmt.Print("Division by zero\n")
    return
  }

  if op == "+" {
    fmt.Println(a + b)
  } else if op == "-" {
    fmt.Println(a - b)
  } else if op == "*" {
    fmt.Println(a * b)
  } else if op == "/" {
    fmt.Println(a / b)
  } else {
    fmt.Print("Invalid operation\n")
  }
}
