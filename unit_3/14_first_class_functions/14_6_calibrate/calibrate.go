package main

import "fmt"

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
	testClosure()
}

func testClosure() {
	var k kelvin = 294.0
	sensor := func() kelvin {
		return k
	}

	fmt.Println(sensor())
	k++
	fmt.Println(sensor())
}

/*
Quick check 14.4
1 What’s another name for an anonymous function in Go?
2 What do closures provide that regular functions don’t?

QC 14.4 answer
1 An anonymous function is also called a function literal in Go.
2 Closures keep references to variables in the surrounding scope
*/
