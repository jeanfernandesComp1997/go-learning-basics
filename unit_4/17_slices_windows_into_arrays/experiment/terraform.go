package main

import "fmt"

type Planets []string

func (planets Planets) terraform() {
	for i, value := range planets {
		planets[i] = "New " + value
	}
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupter", "Saturn", "Uranus", "Neptune",
	}
	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()
	fmt.Println(planets)
}
