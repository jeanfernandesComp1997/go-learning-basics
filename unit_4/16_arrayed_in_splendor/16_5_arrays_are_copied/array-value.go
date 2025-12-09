package main

import "fmt"

func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}

	fmt.Println("Planets terraformed", planets)
}

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	planetsMarkII := planets
	planets[2] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	terraform(planets)
	fmt.Println(planets)
}
