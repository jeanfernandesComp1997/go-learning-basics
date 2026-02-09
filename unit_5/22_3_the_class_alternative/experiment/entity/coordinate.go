package entity

type Coordinate struct {
	D, M, S float64
	H       rune
}

func (coordinate Coordinate) Decimal() float64 {
	sign := 1.0
	switch coordinate.H {
	case 'S', 'W', 's', 'w':
		sign = -1
	}

	return sign * (coordinate.D + coordinate.M/60 + coordinate.S/3600)
}
