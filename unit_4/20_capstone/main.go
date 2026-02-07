package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width      = 80
	height     = 15
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

type Universe [][]bool

func main() {
	universeA, universeB := NewUniverse(), NewUniverse()
	universeA.Seed()

	// for i := 0; i < 500; i++ {
	// 	Step(universeA, universeB)
	// 	universeA.Show()
	// 	time.Sleep(time.Second / 30)
	// 	universeA, universeB = universeB, universeA
	// }

	var grandfather Universe

	for {
		universeA.Show()
		currentSnapshot := universeA.Clone()
		Step(universeA, universeB)

		resetNeeded := false
		reason := ""

		if universeA.Equals(universeB) {
			resetNeeded = true
			reason = "Estabilização detectada!"
		} else if grandfather != nil && grandfather.Equals(universeB) {
			resetNeeded = true
			reason = "Oscilação detectada!"
		}

		if resetNeeded {
			fmt.Printf("\n%s Reiniciando em 2 segundos...", reason)
			time.Sleep(time.Second * 2)
			universeA.Clear()
			universeA.Seed()
			grandfather = nil
			continue
		}

		grandfather = currentSnapshot
		universeA, universeB = universeB, universeA
		time.Sleep(time.Second / 30)
	}
}

func NewUniverse() Universe {
	universe := make(Universe, height)

	for index := range universe {
		universe[index] = make([]bool, width)
	}

	return universe
}

func (universe Universe) String() string {
	buf := make([]byte, 0, (width+1)*height*10)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if universe[y][x] {
				buf = append(buf, colorGreen...)
				buf = append(buf, '*')
				buf = append(buf, colorReset...)
			} else {
				buf = append(buf, ' ')
			}
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

func (universe Universe) Show() {
	fmt.Print("\033[H\033[2J", universe.String())
}

func (universe Universe) Set(x, y int, b bool) {
	universe[y][x] = b
}

func (universe Universe) Seed() {
	aliveCount := int((width * height) * 0.25)

	for aliveCount > 0 {
		x := rand.Intn(width)
		y := rand.Intn(height)

		if !universe[y][x] {
			universe.Set(x, y, true)
			aliveCount--
		}
	}
}

func (universe Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return universe[y][x]
}

func (universe Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && universe.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

func (universe Universe) Next(x, y int) bool {
	neighbors := universe.Neighbors(x, y)
	return neighbors == 3 || neighbors == 2 && universe.Alive(x, y)
}

func Step(universeA, universeB Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			universeB.Set(x, y, universeA.Next(x, y))
		}
	}
}

func (universe Universe) Equals(otherUniverse Universe) bool {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if universe[y][x] != otherUniverse[y][x] {
				return false
			}
		}
	}
	return true
}

func (universe Universe) Clone() Universe {
	newUniverse := NewUniverse()
	for y := 0; y < height; y++ {
		copy(newUniverse[y], universe[y])
	}
	return newUniverse
}

func (universe Universe) Clear() {
	for y := range universe {
		for x := range universe[y] {
			universe[y][x] = false
		}
	}
}
