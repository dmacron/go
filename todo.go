package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Todo CLI")
	scanner := bufio.NewScanner(os.Stdin)

	var tasks []string

	for {
		fmt.Print("Enter command (add/list/exit): ")
		scanner.Scan()
		command := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch command {
		case "add":
			fmt.Println("Enter the task:")
			scanner.Scan()
			task := strings.TrimSpace(scanner.Text())
			tasks = append(tasks, task)
			fmt.Println("Task added.")
		case "list":
			fmt.Println("Your tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		case "exit":
			fmt.Println("Exiting Todo CLI. Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Please use add, list, or exit.")
		}
	}

}
