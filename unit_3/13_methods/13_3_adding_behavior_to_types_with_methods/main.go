package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	var k kelvin = 294.0
	fmt.Print(k, "° K is ", k.celsius(), "° C\n")

	var f fahrenheit = 85.0
	fmt.Print(f, "° F is ", f.celsius(), "° C\n")
}
