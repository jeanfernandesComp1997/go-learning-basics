package main

import "fmt"

type report struct {
	sol         int
	temperature temperature
	location    location
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
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v° C\n", report.temperature.high)
	fmt.Printf("average %v° C\n", t.average())
	fmt.Printf("average %v° C\n", report.average())
}
