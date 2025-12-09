package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1
	for {
		var guess int
		fmt.Println("Guess a number between 1 - 100:")
		fmt.Scanln(&guess)
		if guess < secret {
			fmt.Println("Too low!")
		} else if guess > secret {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Congratulations! You guessed it!")
			break
		}
	}
}
