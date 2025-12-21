package main

import "fmt"

func main() {

	strings := []string{}
	lastCap := cap(strings)

	for i := 0; i < 10000; i++ {
		strings = append(strings, "An element")
		if currentCap := cap(strings); currentCap != lastCap {
			fmt.Println(currentCap)
			lastCap = currentCap
		}
	}
}
