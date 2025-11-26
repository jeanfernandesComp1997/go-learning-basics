package main

import "fmt"

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 1.8) + 32.0)
}

func (c celsius) kelvin() kelvin {
	c += 273.15
	return kelvin(c)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5 / 9)
}

func (f fahrenheit) kelvin() kelvin {
	return kelvin((f-32)*5/9 + 273.15)
}

type kelvin float64

func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit((k-273.15)*1.8 + 32.0)
}

func (k kelvin) celsius() celsius {
	k -= 273.15
	return celsius(k)
}

func main() {
	var c celsius = 28.0
	fmt.Print(c, "° C is ", c.fahrenheit(), "° F\n")
	fmt.Print(c, "° C is ", c.kelvin(), "° K\n")

	var f fahrenheit = 85.0
	fmt.Print(f, "° F is ", f.celsius(), "° C\n")
	fmt.Print(f, "° F is ", f.kelvin(), "° K\n")

	var k kelvin = 300.0
	fmt.Print(k, "° K is ", k.celsius(), "° C\n")
	fmt.Print(k, "° K is ", k.fahrenheit(), "° F\n")
}
