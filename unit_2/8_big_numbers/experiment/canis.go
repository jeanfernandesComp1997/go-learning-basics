package main

import (
	"fmt"
	"math/big"
)

func main() {
	solutionWithConstants()
	solutionWithBig()
}

func solutionWithConstants() {
	const distance = 236_000_000_000_000_000
	const lightSpeed = 299_792
	const secondsPerYear = 31_557_600

	const distanceInLightYears = (distance / (lightSpeed * secondsPerYear))

	fmt.Println("Canis Major Dwarf is", distanceInLightYears, "light years.")
}

func solutionWithBig() {
	var (
		lightSpeed     = big.NewInt(299_792)
		secondsPerYear = big.NewInt(31_557_600)
	)

	kmInLightYear := new(big.Int).Mul(lightSpeed, secondsPerYear)
	distance := big.NewInt(236_000_000_000_000_000)

	distanceInLightYears := new(big.Int)
	distanceInLightYears.Div(distance, kmInLightYear)

	fmt.Println("Canis Major Dwarf is", distanceInLightYears, "light years.")
}
