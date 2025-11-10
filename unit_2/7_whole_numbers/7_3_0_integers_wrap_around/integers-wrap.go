package main

import (
	"fmt"
	"math"
)

func main() {
	var red uint8 = 255
	red++
	fmt.Println(red)

	red = 0
	red--
	fmt.Println(red)

	var number int8 = 127
	number++
	fmt.Println(number)

	var number2 uint16 = math.MaxUint16
	number2++
	fmt.Println(number2)
}
