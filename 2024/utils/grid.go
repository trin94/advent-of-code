package utils

import "fmt"

type Grid struct {
	backend [][]byte
	Columns int
	Rows    int
}

func NewGrid(lines []string) Grid {
	backend := make([][]byte, len(lines))
	for i, line := range lines {
		backend[i] = []byte(line)
	}

	return Grid{
		backend: backend,
		Columns: len(lines[0]),
		Rows:    len(lines),
	}
}

func (g *Grid) Lines() []string {
	lines := make([]string, g.Rows)
	for i, line := range g.backend {
		lines[i] = string(line)
	}
	return lines
}

func (g *Grid) CharAt(row, col int) rune {
	if g.Contains(row, col) {
		return rune(g.backend[row][col])
	} else {
		return ' '
	}
}

func (g *Grid) Contains(row int, col int) bool {
	return row >= 0 && row < g.Rows && col >= 0 && col < g.Columns
}

func (g *Grid) CharAtCoordinate(coordinate Coordinate) rune {
	return g.CharAt(coordinate.Row, coordinate.Col)
}

func (g *Grid) Debug() {
	for _, line := range g.backend {
		fmt.Println(string(line))
	}
}

func (g *Grid) FindCoordinateWithChar(char rune) Coordinate {
	for r := range g.Rows {
		for c := range g.Columns {
			if g.CharAt(r, c) == char {
				return Coordinate{c, r}
			}
		}
	}
	panic(fmt.Sprintf("Coordinate with char '%v' not found\n", char))
}
