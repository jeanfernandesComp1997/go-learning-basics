package main

import (
	"fmt"
	"math"
)

type sol int

func (sol sol) days(sol2 sol) int {
	days := int(sol2 - sol)
	if days < 0 {
		days = -days
	}
	return days
}

type report struct {
	sol
	temperature
	location
}

func (report report) days(s2 sol) int {
	return report.sol.days(s2)
}

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

func (location location) days(location2 location) int {
	marsRadius := 3389.5
	curiosityMetersPerDay := 200.0

	s1, c1 := math.Sincos(rad(location.lat))
	s2, c2 := math.Sincos(rad(location2.lat))
	clong := math.Cos(rad(location.long - location2.long))

	distanceInMeters := (marsRadius * math.Acos(s1*s2+c1*c2*clong)) * 1000

	return int(math.Ceil(distanceInMeters / curiosityMetersPerDay))
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (temperature temperature) average() celsius {
	return (temperature.high + temperature.low) / 2
}

func (report report) average() celsius {
	return report.temperature.average()
}

func main() {
	r := report{
		sol:         15,
		location:    location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}

	fmt.Printf("average %v° C\n", r.average())
	fmt.Printf("average %v° C\n", r.temperature.average())
	fmt.Printf("%v° C\n", r.high)
	r.high = 32
	fmt.Printf("%v° C\n", r.temperature.high)

	r2 := report{sol: 15}

	fmt.Println(r2.sol.days(1446))
	fmt.Println(r2.days(1446))

	location1 := location{-4.5895, 137.4417}
	location2 := location{-4.5895, 142.4417}

	fmt.Printf("Curiosity spend %v days to go from location1 to location2", location1.days(location2))
}
