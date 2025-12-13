package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	records := make([]map[string]string, 0)
	for {
		fmt.Println("Choose(add/list/search/delete/exit):")
		scanner.Scan()
		choice := strings.ToLower(strings.TrimSpace(scanner.Text()))
		switch choice {
		case "add":
			record := make(map[string]string)
			fmt.Println("Enter Name:")
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())
			record["name"] = name
			fmt.Println("Enter Age:")
			scanner.Scan()
			age := strings.TrimSpace(scanner.Text())
			record["age"] = age
			fmt.Println("Enter Grade:")
			scanner.Scan()
			grade := strings.TrimSpace(scanner.Text())
			record["grade"] = grade
			records = append(records, record)
			fmt.Println("Record added.")
		case "list":
			fmt.Println("Student Records:")
			for i, record := range records {
				fmt.Println("Record", i+1, ":", record)
			}
		case "search":
			fmt.Println("Enter Name to search:")
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())
			found := false
			for _, record := range records {
				if record["name"] == name {
					fmt.Println("Record found:", record)
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Record not found.")
			}
		case "delete":
			fmt.Println("Enter Name to delete:")
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())
			index := -1
			for i, record := range records {
				if record["name"] == name {
					index = i
					break
				}
			}
			if index != -1 {
				records = append(records[:index], records[index+1:]...)
				fmt.Println("Record deleted.")
			} else {
				fmt.Println("Record not found.")
			}
		case "exit":
			fmt.Println("Exiting Student Record Manager. Goodbye!")
			return
		}
	}
}
