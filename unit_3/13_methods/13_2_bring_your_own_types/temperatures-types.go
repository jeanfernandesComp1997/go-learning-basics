package main

import "fmt"

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Print(k, "° K is ", c, "° C\n")

	var c2 celsius = 127.0
	k2 := celsiusToKelvin(c2)
	fmt.Print(c2, "° C is ", k2, "° K")
}
