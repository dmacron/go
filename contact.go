package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	contacts := make(map[string]string)

	fmt.Println("=== Contact Book ===")

	for {
		fmt.Print("Choose (add/search/list/delete/exit): ")
		scanner.Scan()
		cmd := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch cmd {
		case "add":
			fmt.Println("Enter name:")
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())
			fmt.Println("Enter phone number:")
			scanner.Scan()
			phone := strings.TrimSpace(scanner.Text())
			contacts[name] = phone
			fmt.Println("Contact added.")
		case "search":
			fmt.Println("Enter name to search:")
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())
			phone, found := contacts[name]
			if found {
				fmt.Printf("Found: %s - %s\n", name, phone)
			} else {
				fmt.Println("Contact not found.")
			}
		case "list":
			fmt.Println("Contacts:")
			for name, phone := range contacts {
				fmt.Printf("%s - %s\n", name, phone)
			}
		case "delete":
			fmt.Println("Enter name to delete:")
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())
			delete(contacts, name)
			fmt.Println("Contact deleted.")
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
