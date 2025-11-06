package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("----- THE GUESS GAME -----")

	var targetNumber = 7

	fmt.Println("The number that the computer have to guess is:", targetNumber)

	var guessAttempt = 0
	var guess int

	for {
		guessAttempt++
		guess = rand.Intn(100) + 1

		fmt.Printf("You choose the number %v and is your %v attempt!\n", guess, guessAttempt)

		if guess > targetNumber {
			fmt.Println("Is too big!")
		} else if guess < targetNumber {
			fmt.Println("Is too small!")
		} else {
			fmt.Printf("Unbelievable! You guess the number at the %v attempt!\n", guessAttempt)
			break
		}
	}
}
