package utils

import (
	"math"
)

type Coordinate struct {
	Col, Row int
}

func ManhattanDistanceBetween(c1, c2 Coordinate) int {
	r := int(math.Abs(float64(c2.Row - c1.Row)))
	c := int(math.Abs(float64(c2.Col - c1.Col)))
	return r + c
}
