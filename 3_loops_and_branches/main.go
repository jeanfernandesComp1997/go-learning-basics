package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// var walkOutside = true
	// var takeTheBluePill = false
	// var wearShades = true

	fmt.Println("You find yourself in a dimly lit cavern.")

	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)

	var containsRead = strings.Contains(command, "read")
	fmt.Println("Can I read the sign?", containsRead)

	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")

	var age = 41
	var minor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	fmt.Println("apple" > "banana")

	branching()

	getRoomDescription()

	leap()

	torch()

	branchWithSwitch()

	branchWithSwitchWithFallthrough()

	repetitionWithLoops()

	loop2()
}

func branching() {
	var command = "go east"

	if command == "go east" {
		fmt.Println("You head further up the montain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didn't quite get that.")
	}
}

func getRoomDescription() {
	var room = "cave"

	if room == "cave" {
		fmt.Println("You find yourself in a dimly lit cavern.")
	} else if room == "entrance" {
		fmt.Println("here is a cavern entrance here and a path to the east.")
	} else if room == "mountain" {
		fmt.Println("There's a cliff here. A path leads west down the mountain.")
	} else {
		fmt.Println("Everything is white.")
	}
}

func leap() {
	fmt.Println("The year is 2100, should you leap?")

	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}
}

func torch() {
	var haveTorch = true
	var litTorch = false

	if !haveTorch || !litTorch {
		fmt.Println("Nothing to see here.")
	}
}

func branchWithSwitch() {
	fmt.Println("There is a cavern entrance here and a path to the east.")
	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You head further up the montain.")
	case "go inside":
		fmt.Println("You enter the cave where you live out the rest of your life.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}
}

func branchWithSwitchWithFallthrough() {
	var room = "lake"

	switch {
	case room == "chave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}

	switch room {
	case "chave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case "underwater":
		fmt.Println("The water is freezing cold.")
	}
}

func repetitionWithLoops() {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff")
}

func loop2() {
	var degrees = 0

	for {
		fmt.Println(degrees)

		degrees++
		if degrees > 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}
