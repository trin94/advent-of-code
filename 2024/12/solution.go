package main

import (
	"os"
	"strings"
)

type Coordinate struct {
	col, row int
}

type FloatCoordinate struct {
	col, row float64
}

type Set[E comparable] map[E]struct{}

func NewSet[E comparable]() Set[E] {
	return Set[E]{}
}

func (s Set[E]) Add(e E) {
	s[e] = struct{}{}
}

func (s Set[E]) Contains(e E) bool {
	_, ok := s[e]
	return ok
}

func (s Set[E]) Remove(e E) {
	delete(s, e)
}

func (s Set[E]) GetNext() (value E, valid bool) {
	for key := range s {
		return key, true
	}
	var zeroValue E
	return zeroValue, false
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

func (g Grid) charAt(row int, col int) string {
	if row >= g.rows || row < 0 {
		return ""
	}
	if col >= g.columns || col < 0 {
		return ""
	}
	return string(g.lines[row][col])
}

func (g Grid) charAtCoordinate(coordinate Coordinate) string {
	return g.charAt(coordinate.row, coordinate.col)
}

func solvePuzzle1(path string) int {
	lines := readLinesFrom(path)
	grid := NewGrid(lines)

	unvisited := NewSet[Coordinate]()
	for r := range grid.rows {
		for c := range grid.columns {
			unvisited.Add(Coordinate{c, r})
		}
	}

	sum := 0

	coordinate, haveMoreUnvisitedFields := unvisited.GetNext()

	for haveMoreUnvisitedFields {

		currentArea := NewSet[Coordinate]()
		currentArea.Add(coordinate)

		currentAreaUnvisited := NewSet[Coordinate]()
		currentAreaUnvisited.Add(coordinate)

		haveMoreFieldsInArea := true
		perimeter := 0

		for haveMoreFieldsInArea {

			perimeter += visit(&grid, coordinate, &currentArea, &currentAreaUnvisited)

			unvisited.Remove(coordinate)
			currentAreaUnvisited.Remove(coordinate)

			coordinate, haveMoreFieldsInArea = currentAreaUnvisited.GetNext()
		}

		sum += perimeter * len(currentArea)

		coordinate, haveMoreUnvisitedFields = unvisited.GetNext()
	}

	return sum
}

func visit(grid *Grid, coordinate Coordinate, currentArea *Set[Coordinate], currentAreaUnvisited *Set[Coordinate]) int {

	char := grid.charAtCoordinate(coordinate)
	perimeterIncrease := 0

	north := Coordinate{coordinate.col, coordinate.row - 1}
	east := Coordinate{coordinate.col + 1, coordinate.row}
	south := Coordinate{coordinate.col, coordinate.row + 1}
	west := Coordinate{coordinate.col - 1, coordinate.row}

	for _, adjacent := range []Coordinate{north, east, south, west} {
		if grid.charAtCoordinate(adjacent) == char {
			if !currentArea.Contains(adjacent) {
				currentArea.Add(adjacent)
				currentAreaUnvisited.Add(adjacent)
			}
		} else {
			perimeterIncrease++
		}
	}

	return perimeterIncrease
}

func solvePuzzle2(path string) int {
	lines := readLinesFrom(path)
	grid := NewGrid(lines)

	unvisited := NewSet[Coordinate]()
	for r := range grid.rows {
		for c := range grid.columns {
			unvisited.Add(Coordinate{c, r})
		}
	}

	sum := 0

	coordinate, haveMoreUnvisitedFields := unvisited.GetNext()

	for haveMoreUnvisitedFields {

		currentArea := NewSet[Coordinate]()
		currentArea.Add(coordinate)

		currentAreaUnvisited := NewSet[Coordinate]()
		currentAreaUnvisited.Add(coordinate)

		haveMoreFieldsInArea := true

		for haveMoreFieldsInArea {

			visit(&grid, coordinate, &currentArea, &currentAreaUnvisited)

			unvisited.Remove(coordinate)
			currentAreaUnvisited.Remove(coordinate)

			coordinate, haveMoreFieldsInArea = currentAreaUnvisited.GetNext()

		}

		sum += countCornersOfShape(&currentArea) * len(currentArea)

		coordinate, haveMoreUnvisitedFields = unvisited.GetNext()
	}

	return sum
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func countCornersOfShape(currentArea *Set[Coordinate]) int {
	// also had the idea of counting corners but struggled to implement it :|
	// https://www.youtube.com/watch?v=KXwKGWSQvS0 @ thank you, dude :)

	cornerCount := 0
	possibleCorners := NewSet[FloatCoordinate]()

	for coordinate := range *currentArea {
		row := float64(coordinate.row)
		col := float64(coordinate.col)
		possibleCorners.Add(FloatCoordinate{col - 0.5, row - 0.5})
		possibleCorners.Add(FloatCoordinate{col - 0.5, row + 0.5})
		possibleCorners.Add(FloatCoordinate{col + 0.5, row + 0.5})
		possibleCorners.Add(FloatCoordinate{col + 0.5, row - 0.5})
	}

	cornersOpposing1 := []bool{true, false, true, false}
	cornersOpposing2 := []bool{false, true, false, true}

	for possibleCorner := range possibleCorners {
		possibleCornerRow := possibleCorner.row
		possibleCornerCol := possibleCorner.col

		edges := 0

		values := make([]bool, 4)

		for idx, cf := range []FloatCoordinate{
			{possibleCornerCol - 0.5, possibleCornerRow - 0.5},
			{possibleCornerCol - 0.5, possibleCornerRow + 0.5},
			{possibleCornerCol + 0.5, possibleCornerRow + 0.5},
			{possibleCornerCol + 0.5, possibleCornerRow - 0.5},
		} {

			cfr := cf.row
			cfc := cf.col

			if currentArea.Contains(Coordinate{int(cfc), int(cfr)}) {
				values[idx] = true
				edges++
			}
		}

		switch {
		case edges == 1 || edges == 3:
			cornerCount++
		case edges == 2:
			if compareSlices(values, cornersOpposing1) || compareSlices(values, cornersOpposing2) {
				cornerCount += 2
			}
		}
	}

	return cornerCount
}

func compareSlices(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
