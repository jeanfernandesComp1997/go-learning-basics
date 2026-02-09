package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

func (coordinate coordinate) decimal() float64 {
	sign := 1.0
	switch coordinate.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}

	return sign * (coordinate.d + coordinate.m/60 + coordinate.s/3600)
}

type location struct {
	lat, long float64
}

func newLocation(lat, long coordinate) location {
	return location{lat: lat.decimal(), long: long.decimal()}
}

func main() {
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	fmt.Println(curiosity)
}
