package main

import (
	"fmt"
	"math/rand"
	"time"
)

const sunrise, sunset = 5, 18

type animal interface {
	move() string
	eat() string
}

type plankton struct {
	name string
}

func (p plankton) String() string {
	return p.name
}

func (p plankton) move() string {
	switch rand.Intn(2) {
	case 0:
		return "photosynthesis"
	default:
		return "adrift"
	}
}

func (p plankton) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "solar light"
	default:
		return "nitrogen and phosphorus"
	}
}

type bee struct {
	name string
}

func (b bee) String() string {
	return b.name
}

func (b bee) move() string {
	switch rand.Intn(2) {
	case 0:
		return "fly around"
	default:
		return "pollinate around"
	}
}

func (b bee) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "nectar and pollen"
	default:
		return "water"
	}
}

type bull struct {
	name string
}

func (b bull) String() string {
	return b.name
}

func (b bull) move() string {
	switch rand.Intn(2) {
	case 0:
		return "walk around"
	default:
		return "run around"
	}
}

func (b bull) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "corn"
	default:
		return "pastures"
	}
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v.\n", a, a.move())
	default:
		fmt.Printf("%v eats the %v.\n", a, a.eat())
	}
}

func main() {
	animals := []animal{
		bee{name: "Normal Bee"},
		plankton{name: "Plankton"},
		bull{name: "Red Bull"},
	}

	var sol, hour int

	for {
		fmt.Printf("%2d:00 ", hour)

		if hour < sunrise || hour >= sunset {
			fmt.Println("The animals are sleeping")
		} else {
			i := rand.Intn(len(animals))
			step(animals[i])
		}

		time.Sleep(500 * time.Millisecond)

		hour++

		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3 {
				break
			}
		}
	}
}
