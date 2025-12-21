package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestrial := planets[0:4:4]
	fmt.Println(terrestrial)

	worlds := append(terrestrial, "Ceres")
	fmt.Println(worlds)
	fmt.Println(terrestrial)
	fmt.Println(planets)

	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	fmt.Println(planets)
}
