package main

import (
	"fmt"
	"math"
)

func main() {
	piggyBank := 0.1
	piggyBank += 0.2

	fmt.Println(piggyBank == 0.3)

	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)

	wallet := 0.0

	for i := 0; i < 11; i++ {
		wallet += 0.10
	}

	fmt.Println("Wallet balance is", wallet)
}
