package main

import (
	"fmt"
	"math/rand"
)

const distance = 62_100_000
const minSpaceshipSpeed = 16
const maxSpaceshipSpeed = 30
const roundTripTripType = "Round-trip"
const oneWayTripType = "One-way"
const earthDayInHours = 24
const secondsInOneHour = 3600
const speedByPriceFactor = 20

func main() {
	fmt.Printf("%-16v %-4v %-11v %v\n", "Spaceline", "Days", "Trip type", "Price in MM")
	fmt.Println("=============================================")

	for i := 0; i < 10; i++ {
		var spaceline string
		switch randomSpacelineValue := rand.Intn(3); randomSpacelineValue {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "Virgin Galactic"
		default:
			spaceline = "SpaceX"
		}

		var tripType string
		if rand.Intn(2) == 0 {
			tripType = oneWayTripType
		} else {
			tripType = roundTripTripType
		}

		tripSpeed := rand.Intn(maxSpaceshipSpeed-minSpaceshipSpeed+1) + minSpaceshipSpeed
		tripDaysDuration := (distance / (tripSpeed * secondsInOneHour)) / earthDayInHours
		if tripType == roundTripTripType {
			tripDaysDuration *= 2
		}

		tripPrice := tripSpeed + speedByPriceFactor
		if tripType == roundTripTripType {
			tripPrice *= 2
		}

		fmt.Printf("%-16v %4v %-11v $ %9v\n", spaceline, tripDaysDuration, tripType, tripPrice)
	}
}
