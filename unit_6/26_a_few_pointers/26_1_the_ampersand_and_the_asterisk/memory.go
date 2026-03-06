package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer)

	address := &answer
	fmt.Println(*address)
	fmt.Printf("address is a %T\n", address)

	brazil := "Brazil"
	var home *string
	fmt.Printf("home is a %T\n", home)

	home = &brazil
	fmt.Println(*home)
}
