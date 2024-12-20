package utils

import "fmt"

type Grid struct {
	lines   []string
	Columns int
	Rows    int
}

func NewGrid(lines []string) Grid {
	return Grid{
		lines:   lines,
		Columns: len(lines[0]),
		Rows:    len(lines),
	}
}

func (g *Grid) Lines() []string {
	return g.lines
}

func (g *Grid) CharAt(row int, col int) rune {
	if g.Contains(row, col) {
		return rune(g.lines[row][col])
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
	for _, line := range g.lines {
		fmt.Println(line)
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
