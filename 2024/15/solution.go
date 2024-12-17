package main

import (
	"fmt"
	"strings"
	"trin94/aoc/2024/inputs"
)

type Coordinate struct {
	row, col int
}

type Grid struct {
	lines   []string
	columns int
	rows    int
}

type Set[E comparable] map[E]struct{}

func NewSet[E comparable]() Set[E] {
	return Set[E]{}
}

func (s Set[E]) Add(v E) {
	s[v] = struct{}{}
}

func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
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

func (g *Grid) setCharAtCoordinate(character rune, coordinate Coordinate) {
	if !g.contains(coordinate.row, coordinate.col) {
		panic("Attempt to move character out of the gird :|")
	}
	lines := g.lines
	out := []rune(lines[coordinate.row])
	out[coordinate.col] = character
	lines[coordinate.row] = string(out)
}

func (g *Grid) robotPosition() Coordinate {
	for row := range g.rows {
		for col := range g.columns {
			if g.charAt(row, col) == '@' {
				return Coordinate{row, col}
			}
		}
	}
	panic("Could not find robot position")
}

func (g *Grid) move(robot Coordinate, direction rune) (newRobotCoordinate Coordinate) {
	var deltaX, deltaY int
	switch direction {
	case '^':
		deltaX = 0
		deltaY = -1
	case 'v':
		deltaX = 0
		deltaY = 1
	case '>':
		deltaX = 1
		deltaY = 0
	case '<':
		deltaX = -1
		deltaY = 0
	}
	steps := 1

	// determine movement
loop:
	for {
		nextCoordinate := Coordinate{
			row: robot.row + deltaY*steps,
			col: robot.col + deltaX*steps,
		}
		nextCharacter := g.charAtCoordinate(nextCoordinate)
		switch nextCharacter {
		case '#':
			steps = 0
			break loop
		case 'O':
			steps++
		case ']':
			steps++
		case '[':
			steps++
		case '.':
			break loop
		default:
			errorMessage := fmt.Sprintf("Unexpected character '%v' at (%d|%d)", nextCharacter, nextCoordinate.row, nextCoordinate.col)
			panic(errorMessage)
		}
	}

	// shift items
	for s := steps; s > 0; s-- {
		c1 := Coordinate{
			row: robot.row + deltaY*(s-1),
			col: robot.col + deltaX*(s-1),
		}
		c2 := Coordinate{
			row: robot.row + deltaY*s,
			col: robot.col + deltaX*s,
		}
		g.swap(c1, c2)
	}

	if steps == 0 {
		return robot
	}

	return Coordinate{robot.row + deltaY, robot.col + deltaX}
}

func (g *Grid) swap(c1, c2 Coordinate) {
	temp := g.charAtCoordinate(c2)
	g.setCharAtCoordinate(g.charAtCoordinate(c1), c2)
	g.setCharAtCoordinate(temp, c1)
}

func (g *Grid) moveScaled(robot Coordinate, direction rune) Coordinate {
	if direction == '<' || direction == '>' {
		return g.move(robot, direction)
	}

	var deltaY int
	switch direction {
	case '^':
		deltaY = -1
	case 'v':
		deltaY = 1
	}

	coordinates, canMove := g.determineCoordinatesToMove(robot, deltaY)
	if !canMove {
		return robot
	}

	var newRobotCoordinate Coordinate
	for i := len(coordinates) - 1; i >= 0; i-- {
		c1 := coordinates[i]
		c2 := Coordinate{
			row: c1.row + deltaY,
			col: c1.col,
		}
		g.swap(c1, c2)

		if i == 0 {
			newRobotCoordinate = c2
		}
	}

	return newRobotCoordinate
}

func (g *Grid) determineCoordinatesToMove(coordinate Coordinate, deltaY int) ([]Coordinate, bool) {
	result := make([]Coordinate, 0)

	currentCoordinates := make([]Coordinate, 0)
	currentCoordinates = append(currentCoordinates, coordinate)

	finished := false

	for !finished {
		nextCoordinates := make([]Coordinate, 0)
		nextCoordinatesSet := NewSet[Coordinate]()

		for _, current := range currentCoordinates {
			next := Coordinate{current.row + deltaY, current.col}

			switch g.charAtCoordinate(next) {
			case '#':
				return result, false
			case '[':
				if !nextCoordinatesSet.Contains(next) {
					nextCoordinatesSet.Add(next)
					nextCoordinates = append(nextCoordinates, next)
				}
				neighbor := Coordinate{current.row + deltaY, current.col + 1}
				if !nextCoordinatesSet.Contains(neighbor) {
					nextCoordinatesSet.Add(neighbor)
					nextCoordinates = append(nextCoordinates, neighbor)
				}
			case ']':
				if !nextCoordinatesSet.Contains(next) {
					nextCoordinatesSet.Add(next)
					nextCoordinates = append(nextCoordinates, next)
				}
				neighbor := Coordinate{current.row + deltaY, current.col - 1}
				if !nextCoordinatesSet.Contains(neighbor) {
					nextCoordinatesSet.Add(neighbor)
					nextCoordinates = append(nextCoordinates, neighbor)
				}
			}

			result = append(result, current)
		}
		finished = len(nextCoordinates) == 0
		currentCoordinates = nextCoordinates
	}

	return result, true
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	grid, robotMovements := parseWarehouse(lines)
	robot := grid.robotPosition()

	for _, movement := range robotMovements {
		robot = grid.move(robot, movement)
	}

	gpsSum := 0

	for r := range grid.rows {
		for c := range grid.columns {
			if grid.charAt(r, c) == 'O' {
				gpsSum += 100*r + c
			}
		}
	}

	return gpsSum
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	grid, robotMovements := parseWarehouse(lines)
	grid = scaleUpWarehouse(grid.lines)
	robot := grid.robotPosition()

	for _, movement := range robotMovements {
		robot = grid.moveScaled(robot, movement)
	}

	gpsSum := 0

	for r := range grid.rows {
		for c := range grid.columns {
			if grid.charAt(r, c) == '[' {
				gpsSum += 100*r + c
			}
		}
	}

	return gpsSum
}

func parseWarehouse(lines []string) (Grid, string) {
	gridLines := make([]string, 0)
	movements := ""

	parseGrid := true

	for _, line := range lines {
		if line == "" {
			parseGrid = false
			continue
		}
		if parseGrid {
			gridLines = append(gridLines, line)
		} else {
			movements += strings.TrimSpace(line)
		}
	}

	return NewGrid(gridLines), movements
}

func scaleUpWarehouse(lines []string) Grid {
	newLines := make([]string, len(lines))
	for r, line := range lines {
		newLine := ""
		for _, character := range line {
			switch character {
			case '#':
				newLine += "##"
			case 'O':
				newLine += "[]"
			case '.':
				newLine += ".."
			case '@':
				newLine += "@."
			}
		}
		newLines[r] = newLine
	}
	return NewGrid(newLines)
}
