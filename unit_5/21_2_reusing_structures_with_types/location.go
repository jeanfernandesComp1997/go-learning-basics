package main

import "fmt"

type location struct {
	lat  float64
	long float64
}

func main() {
	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var oportunity location
	oportunity.lat = -1.9462
	oportunity.long = 354.4734

	fmt.Println(spirit, oportunity)

	var curiosity location
	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity)
}
