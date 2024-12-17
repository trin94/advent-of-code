package utils

type Coordinate struct {
	Col, Row int
}

func NewCoordinate(col, row int) Coordinate {
	return Coordinate{col, row}
}
