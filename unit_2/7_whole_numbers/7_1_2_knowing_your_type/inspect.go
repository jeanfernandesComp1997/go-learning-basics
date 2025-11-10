package main

import "fmt"

func main() {
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days)

	text := "Jean Fernandes"
	fmt.Printf("Type %T for %[1]v\n", text)

	wholeNumber := 7
	fmt.Printf("Type %T for %[1]v\n", wholeNumber)

	realNumber := 9.7
	fmt.Printf("Type %T for %[1]v\n", realNumber)

	booleanValue := true
	fmt.Printf("Type %T for %[1]v\n", booleanValue)
}
