package utils

import "fmt"

type Grid struct {
	lines   []string
	columns int
	rows    int
}

func NewGrid(lines []string) Grid {
	return Grid{
		lines:   lines,
		columns: len(lines[0]),
		rows:    len(lines),
	}
}

func (g *Grid) Lines() []string {
	return g.lines
}

func (g *Grid) Columns() int {
	return g.columns
}

func (g *Grid) Rows() int {
	return g.rows
}

func (g *Grid) CharAt(row int, col int) rune {
	if g.Contains(row, col) {
		return rune(g.lines[row][col])
	} else {
		return ' '
	}
}

func (g *Grid) Contains(row int, col int) bool {
	return row >= 0 && row < g.rows && col >= 0 && col < g.columns
}

func (g *Grid) CharAtCoordinate(coordinate Coordinate) rune {
	return g.CharAt(coordinate.Row, coordinate.Col)
}

func (g *Grid) Debug() {
	for _, line := range g.lines {
		fmt.Println(line)
	}
}
