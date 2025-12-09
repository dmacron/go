package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter two numbers:")
	fmt.Scanln(&a, &b)
	fmt.Println("Choose the operation you wanna perform(+,-,*,/,%):")
	var op string
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Print(a + b)
	case "-":
		fmt.Print(a - b)
	case "*":
		fmt.Print(a * b)
	case "/":
		if b == 0 {
			fmt.Print("Division by zero is not allowed")
			return
		}
		fmt.Print(a / b)
	case "%":
		fmt.Print(a % b)
	default:
		fmt.Print("Invalid operation")
	}
}
