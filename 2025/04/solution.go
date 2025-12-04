package main

import (
	"fmt"
	"trin94/aoc/2025/inputs"
)

type Coordinate struct {
	col, row int
}

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

func (g *Grid) charAt(row int, col int) string {
	if row >= g.rows || row < 0 {
		return ""
	}
	if col >= g.columns || col < 0 {
		return ""
	}
	return string(g.lines[row][col])
}

func (g *Grid) remove(row int, col int) {
	if row >= g.rows || row < 0 {
		panic(fmt.Sprintf("row %d is out of range", row))
	}
	if col >= g.columns || col < 0 {
		panic(fmt.Sprintf("column %d is out of range", col))
	}
	line := []byte(g.lines[row])
	line[col] = '.'
	g.lines[row] = string(line)
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	grid := NewGrid(lines)

	rolls := 0

	for row := 0; row < grid.rows; row++ {
		for col := 0; col < grid.columns; col++ {
			if grid.charAt(row, col) == "@" {
				adjacent := countAdjacent(grid, row, col)
				if adjacent < 4 {
					rolls++
				}
			}
		}
	}

	return rolls
}

func countAdjacent(grid Grid, row int, col int) int {
	count := 0
	if grid.charAt(row-1, col-1) == "@" {
		count += 1
	}
	if grid.charAt(row-1, col) == "@" {
		count += 1
	}
	if grid.charAt(row-1, col+1) == "@" {
		count += 1
	}
	if grid.charAt(row, col-1) == "@" {
		count += 1
	}
	if grid.charAt(row, col+1) == "@" {
		count += 1
	}
	if grid.charAt(row+1, col-1) == "@" {
		count += 1
	}
	if grid.charAt(row+1, col) == "@" {
		count += 1
	}
	if grid.charAt(row+1, col+1) == "@" {
		count += 1
	}
	return count
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	grid := NewGrid(lines)

	rolls := 0

	for {
		accessed := make([]Coordinate, 0)

		for row := 0; row < grid.rows; row++ {
			for col := 0; col < grid.columns; col++ {
				if grid.charAt(row, col) == "@" {
					adjacent := countAdjacent(grid, row, col)
					if adjacent < 4 {
						accessed = append(accessed, Coordinate{col, row})
						rolls++
					}
				}
			}
		}

		if len(accessed) == 0 {
			break
		}

		for _, paper := range accessed {
			grid.remove(paper.row, paper.col)
		}
	}

	return rolls
}
