package utils

type Coordinate struct {
	Col, Row int
}

func NewCoordinate(col, row int) Coordinate {
	return Coordinate{col, row}
}

func (c Coordinate) North() Coordinate {
	return Coordinate{c.Col, c.Row - 1}
}

func (c Coordinate) South() Coordinate {
	return Coordinate{c.Col, c.Row + 1}
}

func (c Coordinate) East() Coordinate {
	return Coordinate{c.Col + 1, c.Row}
}

func (c Coordinate) West() Coordinate {
	return Coordinate{c.Col - 1, c.Row}
}
