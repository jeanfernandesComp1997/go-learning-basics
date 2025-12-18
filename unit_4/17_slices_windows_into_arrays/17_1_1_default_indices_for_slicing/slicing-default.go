package main

import "fmt"

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

	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]

	fmt.Println(terrestrial, gasGiants, iceGiants)

	allPlanets := planets[:]
	fmt.Println(allPlanets)

	neptune := "Neptune"
	tune := neptune[3:]
	fmt.Println(tune)

	neptune = "Poseidon"
	fmt.Println(tune)

	question := "¿Cómo estás?"
	fmt.Println(question[:6])

	colonizedPlanets := terrestrial[2:]
	fmt.Println(colonizedPlanets)
}
