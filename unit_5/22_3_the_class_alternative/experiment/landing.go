package main

import (
	"experiment/entity"
)

func main() {
	spiritLatCoordinate := entity.Coordinate{D: 14, M: 34, S: 6.2, H: 'S'}
	spiritLongCoordinate := entity.Coordinate{D: 175, M: 28, S: 24.2, H: 'E'}

	spirit := entity.Location{Lat: spiritLatCoordinate.Decimal(), Long: spiritLongCoordinate.Decimal()}
	spirit.Print()
}
