package main

import "fmt"

func printItems(items []string) {
	fmt.Println(items)
}

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	printItems(dwarfs)

	dwarfsArray := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfsSlice := dwarfsArray[:]
	printItems(dwarfsSlice)

	fmt.Printf("%T %T", dwarfsArray, dwarfs)
}
