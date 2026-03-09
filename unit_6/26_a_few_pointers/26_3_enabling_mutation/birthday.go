package main

import (
	"fmt"
	"time"
)

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func (p *person) birthday() {
	p.age++
}

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}

func main() {
	jean := person{
		name:       "Jean",
		superpower: "imagination",
		age:        28,
	}

	birthday(&jean)
	fmt.Printf("%+v\n", jean)

	carol := &person{
		name:       "Carol",
		superpower: "intelligence",
		age:        27,
	}

	carol.birthday()
	fmt.Printf("%+v\n", carol)

	gabriel := person{
		name:       "Gabriel",
		age:        13,
		superpower: "super strength",
	}

	gabriel.birthday()
	fmt.Printf("%+v\n", gabriel)

	const layout = "Mon, Jan 2, 2006"
	day := time.Now()
	tomorrow := day.Add(24 * time.Hour)
	fmt.Println(day.Format(layout))
	fmt.Println(tomorrow.Format(layout))

	player := character{name: "Jean"}
	levelUp(&player.stats)
	fmt.Printf("%v\n", player.stats)

	var board [8][8]rune
	reset(&board)
	fmt.Printf("%c", board[0][0])
}
