package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func measureTemperature(samples int, s sensor) {
	for i := 0; i < samples; i++ {
		k := s()
		fmt.Printf("%v K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	measureTemperature(3, fakeSensor)
}

// Quick check
// Rewrite the following function signature to use a function type:
// func drawTable(rows int, getRow func(row int) (string, string))

type getRowFn func(row int) (string, string)

func drawTable(rows int, getRow getRowFn)
