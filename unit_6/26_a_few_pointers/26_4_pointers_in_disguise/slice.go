package main

import "fmt"

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func reclassifyBetterApproach(planets []string) []string {
	return planets[0:8]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets)
	fmt.Println(planets)

	planets2 := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	fmt.Println(reclassifyBetterApproach(planets2))
}
