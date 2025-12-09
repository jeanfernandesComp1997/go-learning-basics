package main

import "fmt"

func main() {
	var planets [8]string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	// planets[3] = "Mars"
	// planets[4] = "Jupiter"
	// planets[5] = "Saturn"
	// planets[6] = "Uranus"
	// planets[7] = "Neptune"

	earth := planets[2]
	fmt.Println(earth)

	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")

	var numbers [10]int
	fmt.Println(numbers[0])
}
