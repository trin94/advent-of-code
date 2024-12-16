package main

import (
	"math"
	"os"
	"strings"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

/// ###

type Set[E comparable] map[E]struct{}

func NewSet[E comparable]() Set[E] {
	return Set[E]{}
}

func (s Set[E]) Add(e E) {
	s[e] = struct{}{}
}

/// ###

type Coordinate struct {
	col, row int
}

/// ###

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

func (g *Grid) charAt(row int, col int) rune {
	if g.contains(row, col) {
		return rune(g.lines[row][col])
	} else {
		return ' '
	}
}

func (g *Grid) contains(row int, col int) bool {
	return row >= 0 && row < g.rows && col >= 0 && col < g.columns
}

func (g *Grid) charAtCoordinate(coordinate Coordinate) rune {
	return g.charAt(coordinate.row, coordinate.col)
}

/// ###

type MazeField struct {
	coordinate     Coordinate
	direction      Direction
	cost           int
	previousFields []*MazeField
}

/// ###

func solvePuzzle1(path string) int {
	lines := readLinesFrom(path)
	grid, start := parseGrid(lines)
	pathsToEnd := traverseMaze(grid, start, false)
	return pathsToEnd[0].cost
}

func solvePuzzle2(path string) int {
	lines := readLinesFrom(path)
	grid, start := parseGrid(lines)
	pathsToEnd := traverseMaze(grid, start, true)
	coordinates := NewSet[Coordinate]()

	for _, pathToEnd := range pathsToEnd {
		for _, field := range pathToEnd.previousFields {
			coordinates.Add(field.coordinate)
		}
	}

	return len(coordinates)
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func parseGrid(lines []string) (grid Grid, start Coordinate) {
	grid = NewGrid(lines)
	for r := range grid.rows {
		for c := range grid.columns {
			if grid.charAt(r, c) == 'S' {
				return grid, Coordinate{c, r}
			}
		}
	}
	panic("Could not find start and end coordinates")
}

func traverseMaze(grid Grid, start Coordinate, allResults bool) []*MazeField {
	startField := &MazeField{start, East, 0, make([]*MazeField, 0)}
	startField.previousFields = append(startField.previousFields, startField)

	visitNext := make([]*MazeField, 0)
	visitNext = append(visitNext, startField)

	minimalCosts := map[Coordinate]int{}

	maxCost := math.MaxInt
	best := make([]*MazeField, 0)

	costPuffer := 0
	if allResults {
		costPuffer = 2*1001 + 1
	}

	for len(visitNext) != 0 {

		current := visitNext[0]
		visitNext = visitNext[1:]

		if current.cost > maxCost {
			continue
		}

		directions := []Direction{North, East, South, West}
		neighbors := []Coordinate{
			{current.coordinate.col, current.coordinate.row - 1},
			{current.coordinate.col + 1, current.coordinate.row},
			{current.coordinate.col, current.coordinate.row + 1},
			{current.coordinate.col - 1, current.coordinate.row},
		}

		for i, neighbor := range neighbors {

			character := grid.charAtCoordinate(neighbor)
			if character != '.' && character != 'S' && character != 'E' {
				continue
			}

			direction := directions[i]
			costDelta := determineCost(current.direction, direction)
			if costDelta == 2002 {
				continue
			}
			costs := current.cost + costDelta

			minimalCostSoFar, found := minimalCosts[neighbor]
			if !found || costs < minimalCostSoFar {
				minimalCosts[neighbor] = costs
			} else if costs >= minimalCostSoFar+costPuffer {
				continue
			}

			mazeField := &MazeField{coordinate: neighbor, direction: direction, cost: costs}
			mazeField.previousFields = append(clone(current.previousFields), mazeField)

			if character == 'E' {
				if maxCost > mazeField.cost {
					maxCost = mazeField.cost
					best = make([]*MazeField, 0)
				}
				best = append(best, mazeField)
				continue
			}

			visitNext = append(visitNext, mazeField)

		}

	}

	return best
}

func determineCost(source Direction, target Direction) int {
	if source != target {
		if (source+2)%4 == target {
			return 2002
		}
		return 1001
	}
	return 1
}

func clone[E comparable](slice []E) []E {
	var dst []E
	return append(dst, slice...)
}
