package utils

import (
	"fmt"
)

type Grid struct {
	backend [][]byte
	columns int
	rows    int
}

func NewGrid(lines []string) Grid {
	backend := make([][]byte, len(lines))
	for i, line := range lines {
		backend[i] = []byte(line)
	}

	return Grid{
		backend: backend,
		columns: len(lines[0]),
		rows:    len(lines),
	}
}

func (g *Grid) Lines() []string {
	lines := make([]string, g.rows)
	for i, line := range g.backend {
		lines[i] = string(line)
	}
	return lines
}

func (g *Grid) Columns() int {
	return g.columns
}

func (g *Grid) Rows() int {
	return g.rows
}

func (g *Grid) CharAt(row, col int) rune {
	if g.Contains(row, col) {
		return rune(g.backend[row][col])
	} else {
		return ' '
	}
}

func (g *Grid) Contains(row, col int) bool {
	return row >= 0 && row < g.rows && col >= 0 && col < g.columns
}

func (g *Grid) CharAtCoordinate(coordinate Coordinate) rune {
	return g.CharAt(coordinate.Row, coordinate.Col)
}

func (g *Grid) Debug() {
	for _, line := range g.backend {
		fmt.Println(string(line))
	}
}

func (g *Grid) SetCharAt(char rune, row, col int) {
	g.backend[row][col] = byte(char)
}
