package main

import "fmt"

func main() {
	age := 28
	marsAge := float64(age)

	marsDays := 687.0
	earthDays := 365.2425
	marsAge *= earthDays / marsDays
	fmt.Println("I am", marsAge, "years old on Mars.")

	fmt.Println(int(earthDays))
}
