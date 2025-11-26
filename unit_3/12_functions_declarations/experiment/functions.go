package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32.0
}

func kelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*1.8 + 32.0
}

func main() {
	kelvin := 233.0
	fmt.Print(kelvin, "° K is ", kelvinToCelsius(kelvin), "° C\n")

	celcius := 28.0
	fmt.Print(celcius, "° C is ", celsiusToFahrenheit(celcius), "° F\n")

	kelvin2 := 0.0
	fmt.Printf("%.2f ° K is %.2f° F\n", kelvin2, kelvinToFahrenheit(kelvin2))
}
