package main

import (
    "fmt"
)

func Add(a int, b int) int {
    return a + b
}

func Subtract(a int, b int) int {
    return a - b
}

func main() {
    fmt.Println("Result:", Add(5, 3))
}
