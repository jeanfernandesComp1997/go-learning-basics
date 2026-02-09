package main

import "fmt"

type location struct {
	lat, long float64
}

func main() {
	oportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(oportunity)

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)

	spirity := location{-14.5684, 175.472636}
	fmt.Println(spirity)

	curiosity := location{-4.5895, 137.4417}
	fmt.Printf("%v\n", curiosity)
	fmt.Printf("%+v\n", curiosity)
}
