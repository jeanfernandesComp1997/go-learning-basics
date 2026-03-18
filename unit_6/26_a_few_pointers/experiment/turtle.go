package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y--
}

func (t *turtle) down() {
	t.y++
}

func (t *turtle) right() {
	t.x++
}

func (t *turtle) left() {
	t.x--
}

func main() {
	leo := turtle{}

	leo.up()
	leo.up()
	leo.left()
	leo.left()
	fmt.Println(leo)

	leo.down()
	leo.down()
	leo.right()
	leo.right()
	fmt.Println(leo)
}
