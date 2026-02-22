package main

import (
	"fmt"
	"math"
)

type location struct {
	name      string
	lat, long float64
}

func (location location) description() string {
	return fmt.Sprintf("Name: %v, Latitude: %.1f°, Longitude: %.1f°", location.name, location.lat, location.long)
}

type world struct {
	radius float64
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (world world) distance(currentLocation, destinationLocation location) float64 {
	s1, c1 := math.Sincos(rad(currentLocation.lat))
	s2, c2 := math.Sincos(rad(destinationLocation.lat))
	clong := math.Cos(rad(currentLocation.long - destinationLocation.long))

	return world.radius * math.Acos(s1*s2+c1*c2*clong)
}

type gps struct {
	currentLocation, destinationLocation location
	world
}

func (gps gps) distance() float64 {
	return gps.world.distance(gps.currentLocation, gps.destinationLocation)
}

func (gps gps) message() string {
	return fmt.Sprintf("The remaining distance is %.1f km to %v", gps.distance(), gps.destinationLocation.description())
}

type rover struct {
	gps
}

func main() {
	mars := world{radius: 3389.5}

	bradBury := location{lat: -4.5895, long: 137.4417}
	elysiumPlanitia := location{lat: 4.5, long: 135.9}

	gps := gps{
		currentLocation:     bradBury,
		destinationLocation: elysiumPlanitia,
		world:               mars,
	}

	curiosityRover := rover{gps: gps}

	fmt.Println(curiosityRover.message())
}
