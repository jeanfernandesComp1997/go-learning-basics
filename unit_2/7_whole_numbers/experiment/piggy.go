package main

import (
	"fmt"
	"math/rand"
)

const TARGET_BALLANCE uint16 = 2000
const CENTS_IN_A_DOLLAR float32 = 100

func main() {
	walletBalance := uint16(0)

	for walletBalance < TARGET_BALLANCE {
		switch coinType := rand.Intn(3); coinType {
		case 0:
			walletBalance += 5
		case 1:
			walletBalance += 10
		default:
			walletBalance += 25
		}

		fmt.Printf("Actual wallet balance: %.2f\n", float32(walletBalance)/CENTS_IN_A_DOLLAR)
	}

	fmt.Printf("Final wallet balance is: %.2f\n", float32(walletBalance)/CENTS_IN_A_DOLLAR)
}
