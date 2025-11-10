package main

import (
	"fmt"
	"math/rand"
)

const TARGET_BALLANCE = 20.0

func main() {
	walletBalance := 0.0

	for walletBalance < TARGET_BALLANCE {
		switch coinType := rand.Intn(3); coinType {
		case 0:
			walletBalance += 0.05
		case 1:
			walletBalance += 0.10
		default:
			walletBalance += 0.25
		}

		fmt.Printf("Actual wallet balance: %.2f\n", walletBalance)
	}

	fmt.Printf("Final wallet balance is: %.2f\n", walletBalance)
}
