package main

import "fmt"

func main() {
	value := "yes"
	var convertedValue bool

	switch value {
	case "yes", "true", "1":
		convertedValue = true
	case "no", "false", "0":
		convertedValue = false
	default:
		fmt.Println("Invalid value to convert to boolean!")
	}

	fmt.Println("Converted value is:", convertedValue)

}
