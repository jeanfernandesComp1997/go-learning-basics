package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

// func (l laser) talk() string {
// 	return strings.Repeat("pew ", int(l))
// }

func (l laser) talk() string {
	return strings.Repeat("toot ", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type crater struct{}

type starship struct {
	laser
}

type rover string

func (r rover) talk() string {
	return string(r)
}

func main() {
	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	shout(martian{})
	shout(laser(2))

	// shout(crater{}) compile error

	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)

	r := rover("whir whir")
	shout(r)
}
