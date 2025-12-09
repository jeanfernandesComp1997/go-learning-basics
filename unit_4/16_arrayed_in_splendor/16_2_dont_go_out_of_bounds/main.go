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
	// planets[8] = "Pluto" // Invalid array index 8 (out of bounds for 8-element array)

	i := 8
	planets[i] = "Pluto"
	pluto := planets[i]
	fmt.Println(pluto)
	// Panic: runtime error: index out of range
}
