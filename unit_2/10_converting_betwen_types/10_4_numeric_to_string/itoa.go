package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown := 10

	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	fmt.Println(str)

	countdown = 9
	str = fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(str)

	newCountdown, err := strconv.Atoi("10")
	if err != nil {
		// oh no, something went wrong
	}
	fmt.Println(newCountdown)
}
