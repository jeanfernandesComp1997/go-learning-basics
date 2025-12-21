package main

import (
	"fmt"
	"slices"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupter", "Saturn", "Uranus", "Neptune",
	}
	// sort.StringSlice(planets).Sort()
	// fmt.Println(planets)
	slices.Sort(planets)
	fmt.Println(planets)
}
