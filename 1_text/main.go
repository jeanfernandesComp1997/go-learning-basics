// My weight loss program
package main

import "fmt"

// main is the function where it all begins
func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(60 * 0.3783)
	fmt.Print(" kg, and I would be ")
	fmt.Print(27 * 365 / 687)
	fmt.Print(" years old.")

	fmt.Println()

	fmt.Printf("My weight on the surface of Mars is %v kg,", 60*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 27*365/687)

	fmt.Printf("My weight on the surface of %v is %v kg.\n", "Earth", 60)

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
