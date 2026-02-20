package main

import "fmt"

type sol int

func (sol sol) days(sol2 sol) int {
	days := int(sol2 - sol)
	if days < 0 {
		days = -days
	}
	return days
}

type report struct {
	sol
	temperature
	location
}

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

func (temperature temperature) average() celsius {
	return (temperature.high + temperature.low) / 2
}

func (report report) average() celsius {
	return report.temperature.average()
}

func main() {
	r := report{
		sol:         15,
		location:    location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}

	fmt.Printf("average %v° C\n", r.average())
	fmt.Printf("average %v° C\n", r.temperature.average())
	fmt.Printf("%v° C\n", r.high)
	r.high = 32
	fmt.Printf("%v° C\n", r.temperature.high)

	r2 := report{sol: 15}

	fmt.Println(r2.sol.days(1446))
	fmt.Println(r2.days(1446))
}
