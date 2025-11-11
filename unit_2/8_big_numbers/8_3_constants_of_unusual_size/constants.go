package main

import "fmt"

func main() {
	fmt.Println("Andromeda Galaxy is", 24000000000000000000/299792/86400, "light days away.")

	const distance = 24000000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400

	// The Go compiler simplifies equations containing constants and literals during compilation.
	const days = distance / lightSpeed / secondsPerDay

	fmt.Println("Andromeda Galaxy is", days, "light days away.")

	// km := distance
	dayss := distance / lightSpeed / secondsPerDay
	fmt.Println("Andromeda Galaxy is", dayss, "light days away.")

	// fmt.Println("Andromeda Galaxy is", distance, "km away.")
}
