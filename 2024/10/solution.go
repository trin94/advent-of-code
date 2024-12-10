package main

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

// ###

type Set[E comparable] map[E]struct{}

func NewSet[E comparable]() Set[E] {
	return Set[E]{}
}

func (s Set[E]) Add(e E) {
	s[e] = struct{}{}
}

func (s Set[E]) AddAll(set Set[E]) {
	for e := range set {
		s[e] = struct{}{}
	}
}

func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}

type Coordinate struct {
	col, row int
}

type Grid struct {
	lines   [][]int
	columns int
	rows    int
}

func NewGrid(lines []string) Grid {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, value := range line {
			if unicode.IsNumber(value) {
				result, _ := strconv.Atoi(string(value))
				row[j] = result
			} else {
				row[j] = -1
			}
		}
		grid[i] = row
	}
	return Grid{
		lines:   grid,
		columns: len(grid[0]),
		rows:    len(lines),
	}
}

func (g Grid) GetTrailheads() []Coordinate {
	result := make([]Coordinate, 0)
	for row := 0; row < g.rows; row++ {
		for col := 0; col < g.columns; col++ {
			if g.lines[row][col] == 0 {
				result = append(result, Coordinate{col, row})
			}
		}
	}
	return result
}

func (g Grid) HeightAt(row, col int) int {
	if row < 0 || row >= g.rows || col < 0 || col >= g.columns {
		return -1
	}
	return g.lines[row][col]
}

func solvePuzzle1(path string) int {
	numbers := readNumbersFrom(path)
	grid := NewGrid(numbers)

	sum := 0

	for _, th := range grid.GetTrailheads() {
		summits := countReachableSummits(&grid, th, 1)
		sum += len(summits)
	}

	return sum
}

func solvePuzzle2(path string) int {
	numbers := readNumbersFrom(path)
	grid := NewGrid(numbers)

	sum := 0

	for _, th := range grid.GetTrailheads() {
		sum += countTrails(&grid, th, 1)
	}

	return sum
}

func readNumbersFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func countReachableSummits(grid *Grid, location Coordinate, nextHeight int) Set[Coordinate] {
	summits := NewSet[Coordinate]()
	if nextHeight == 10 {
		summits.Add(location)
		return summits
	}
	if grid.HeightAt(location.row-1, location.col) == nextHeight {
		north := Coordinate{location.col, location.row - 1}
		more := countReachableSummits(grid, north, nextHeight+1)
		summits.AddAll(more)
	}
	if grid.HeightAt(location.row, location.col+1) == nextHeight {
		east := Coordinate{location.col + 1, location.row}
		more := countReachableSummits(grid, east, nextHeight+1)
		summits.AddAll(more)
	}
	if grid.HeightAt(location.row+1, location.col) == nextHeight {
		south := Coordinate{location.col, location.row + 1}
		more := countReachableSummits(grid, south, nextHeight+1)
		summits.AddAll(more)
	}
	if grid.HeightAt(location.row, location.col-1) == nextHeight {
		west := Coordinate{location.col - 1, location.row}
		more := countReachableSummits(grid, west, nextHeight+1)
		summits.AddAll(more)
	}
	return summits
}

func countTrails(grid *Grid, location Coordinate, nextHeight int) int {
	if nextHeight == 10 {
		return 1
	}
	result := 0
	if grid.HeightAt(location.row-1, location.col) == nextHeight {
		north := Coordinate{location.col, location.row - 1}
		result += countTrails(grid, north, nextHeight+1)
	}
	if grid.HeightAt(location.row, location.col+1) == nextHeight {
		east := Coordinate{location.col + 1, location.row}
		result += countTrails(grid, east, nextHeight+1)
	}
	if grid.HeightAt(location.row+1, location.col) == nextHeight {
		south := Coordinate{location.col, location.row + 1}
		result += countTrails(grid, south, nextHeight+1)
	}
	if grid.HeightAt(location.row, location.col-1) == nextHeight {
		west := Coordinate{location.col - 1, location.row}
		result += countTrails(grid, west, nextHeight+1)
	}
	return result
}
