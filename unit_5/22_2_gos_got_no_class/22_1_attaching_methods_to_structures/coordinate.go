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

func main() {
	// Bradbury Landing: 4º35'22.2" S, 137º26'30.1" E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())
}
