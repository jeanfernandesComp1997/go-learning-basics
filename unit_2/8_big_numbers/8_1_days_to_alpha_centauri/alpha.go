package main

import "fmt"

func main() {
	const lightSpeedInKmByHour = 299792
	const secondsPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha centaury is", distance, "km away.")

	days := distance / lightSpeedInKmByHour / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")

	var distanceMarsAndEarth int64 = 56e6
	var farthestDistanceMarsAndEarth int64 = 401e6
	fmt.Println(distanceMarsAndEarth)
	fmt.Println(farthestDistanceMarsAndEarth)

}
