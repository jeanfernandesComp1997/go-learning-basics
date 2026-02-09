package main

import (
	"experiment/entity"
)

func main() {
	// Spirit
	spirit := entity.NewLocation(
		"Spirit",
		"Columbia Memorial Station",
		entity.Coordinate{D: 14, M: 34, S: 6.2, H: 'S'},
		entity.Coordinate{D: 175, M: 28, S: 21.5, H: 'E'},
	)
	spirit.Print()

	// Opportunity
	opportunity := entity.NewLocation(
		"Opportunity",
		"Challenger Memorial Station",
		entity.Coordinate{D: 1, M: 56, S: 46.3, H: 'S'},
		entity.Coordinate{D: 354, M: 28, S: 24.2, H: 'E'},
	)
	opportunity.Print()

	// Curiosity
	curiosity := entity.NewLocation(
		"Curiosity",
		"Bradbury Landing",
		entity.Coordinate{D: 4, M: 35, S: 22.2, H: 'S'},
		entity.Coordinate{D: 137, M: 26, S: 30.1, H: 'E'},
	)
	curiosity.Print()

	// InSight
	insight := entity.NewLocation(
		"InSight",
		"Elysium Planitia",
		entity.Coordinate{D: 4, M: 30, S: 0.0, H: 'N'},
		entity.Coordinate{D: 135, M: 54, S: 0, H: 'E'},
	)
	insight.Print()
}
