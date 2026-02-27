package main

import (
	"fmt"
	"time"
)

const (
	starDateEpoch      = 1000.0
	hoursPerDay        = 24.0
	solsInAMartianYear = 668
	unknownHour        = 0
)

type starDater interface {
	YearDay() int
	Hour() int
}

func starDate(t starDater) float64 {
	dayOfYear := float64(t.YearDay())
	hourFraction := float64(t.Hour()) / hoursPerDay
	return starDateEpoch + dayOfYear + hourFraction
}

// func starDate(t time.Time) float64 {
// 	dayOfYear := float64(t.YearDay())
// 	hourFraction := float64(t.Hour()) / hoursPerDay
// 	return starDateEpoch + dayOfYear + hourFraction
// }

type sol int

func (s sol) YearDay() int {
	return int(s % solsInAMartianYear)
}

func (s sol) Hour() int {
	return unknownHour
}

func main() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", starDate(day))

	s := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", starDate(s))
}
