// How long does it take to get to Mars?
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const lightSpeedInKmByS = 299792
	var distanceInKm = 56000000

	fmt.Println(distanceInKm/lightSpeedInKmByS, "seconds")

	distanceInKm = 401000000
	fmt.Println(distanceInKm/lightSpeedInKmByS, "seconds")

	var spaceShipSpeedInKmByH = 100800
	var distanceBetweenEarthAndMarsInKm = 96300000

	// var (
	// 	distance = 56000000
	// 	speed    = 100800
	// )

	// var distance, speed = 56000000, 100800

	// const hoursPerDay, minutesPerHour = 24, 60

	const earthDayInHours = 24
	fmt.Println((distanceBetweenEarthAndMarsInKm/spaceShipSpeedInKmByH)/earthDayInHours, "days")

	var weight = 60.0
	weight = weight * 0.3783
	weight *= 0.3783
	weight -= 2

	var age = 27
	age = age + 1
	age += 1
	age++
	age--
	// price /=2

	randomNumbers()
	generateRandomDistanceBetweenEarthAndMarsInKm()
	generateRandomDistanceBetweenEarthAndMarsInKmAtBookWay()
	malacandraDistance()
}

func randomNumbers() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}

func generateRandomDistanceBetweenEarthAndMarsInKm() {
	var min = 56000000
	var max = 401000000

	var distance = rand.Intn(max-min+1) + min
	fmt.Printf("Distance from Earth to Mards right now is %v km\n", distance)
}

func generateRandomDistanceBetweenEarthAndMarsInKmAtBookWay() {
	var distance = rand.Intn(345000001) + 56000000
	fmt.Printf("Distance from Earth to Mards right now is %v km\n", distance)
}

func malacandraDistance() {
	const distance = 56000000
	const arrivalInDays = 28
	const hoursInADay = 24
	var arrivalInHours = arrivalInDays * hoursInADay

	var necessarySpeed = distance / arrivalInHours
	fmt.Printf("The speed must be %v km/h", necessarySpeed)
}
