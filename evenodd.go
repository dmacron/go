package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("It's an even number")
	} else {
		fmt.Println("It's an odd number")
	}
}
