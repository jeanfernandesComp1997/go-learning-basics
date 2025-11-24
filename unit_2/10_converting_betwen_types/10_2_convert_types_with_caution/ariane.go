package main

import (
	"fmt"
	"math"
)

func main() {
	var bh float64 = 32767
	var h = int16(bh)
	fmt.Println(h)

	if bh < math.MinInt16 || bh > math.MaxInt16 {
		// handle out of range value
	}

	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		v8 := uint8(v)
		fmt.Println("Converted:", v8)
	}
}
