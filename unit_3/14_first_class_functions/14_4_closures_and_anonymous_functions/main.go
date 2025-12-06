package main

import "fmt"

var f = func() {
	fmt.Println("Dress up for the masquerade.")
}

func main() {
	f()

	f := func(message string) {
		fmt.Println(message)
	}

	f("Go to the party.")

	func() {
		fmt.Println("Functions anonymous")
	}()
}
