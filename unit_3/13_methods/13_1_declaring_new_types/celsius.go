package main

import "fmt"

func main() {
	type celsius float64

	var temperature celsius = 20
	fmt.Println(temperature)

	const degrees = 20
	var temperature2 celsius = degrees

	temperature2 += 10
	fmt.Println(temperature2)

	// Invalid operation: mismatched types
	// var warmUp float64 = 10
	// temperature2 += warm

	// type fahrenheit float64

	// var c celsius = 20
	// var f fahrenheit = 20

	// Invalid operation: mismatched types
	// if c == f {

	// }

	// c += f
}
