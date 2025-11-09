package main

import "fmt"

func main() {
	third := 1.0 / 3

	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)
	fmt.Printf("%05.2f\n", third)
	fmt.Printf("%7.4f\n", third)
	fmt.Printf("%06.2f\n", third)
}
