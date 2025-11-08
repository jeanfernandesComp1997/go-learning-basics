package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	daysInMonth := 31
	divisibleBy4 := 4
	divisibleBy400 := 400
	maxYear := 3000
	minYear := 2026

	for i := 0; i < 10; i++ {
		year := rand.Intn(maxYear-minYear) + minYear
		month := rand.Intn(12) + 1

		switch month {
		case 2:
			if year%divisibleBy4 == 0 || year%divisibleBy400 == 0 {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		day := rand.Intn(daysInMonth) + 1

		fmt.Println(era, year, month, day)
	}
}
