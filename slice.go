package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	items := make([]string, 0, 25)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Itemlister - Choose an option: 1. Add Item 2. List Items 3.Check Length 4.Check Capacity 5.Lookthrough 6.Delete an item 7.Exit")
		scanner.Scan()
		c := scanner.Text()
		switch c {
		case "1":
			fmt.Println("Enter item to add:")
			scanner.Scan()
			item := scanner.Text()
			items = append(items, item)
			fmt.Println("Item added.")
		case "2":
			fmt.Println("Items:")
			for i, item := range items {
				fmt.Printf("%d. %s\n", i+1, item)
			}
		case "3":
			fmt.Printf("Length: %d\n", len(items))
		case "4":
			fmt.Printf("Capacity: %d\n", cap(items))
		case "5":
			look := items[0 : len(items)-1]
			fmt.Println("Looking through items:", look)
		case "6":
			fmt.Println("Enter index of item to delete:")
			var index int
			fmt.Scanln(&index)
			items = append(items[:index], items[index+1:]...)
			fmt.Println("Item deleted.")
		case "7":
			fmt.Println("Exiting Itemlister. Goodbye!")
			return
		}
	}
}
