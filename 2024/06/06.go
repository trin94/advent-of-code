package p06

import (
	"trin94/aoc/2024/inputs"
)

// ###

type Coordinate struct {
	col, row int
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

// ###

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func parseDirection(value string) Direction {
	switch value {
	case "^":
		return North
	case "v":
		return South
	case ">":
		return East
	case "<":
		return West
	default:
		return None
	}
}

func (d Direction) rotate() Direction {
	switch d {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	default:
		panic("unreachable")
	}
}

// ###

type Marker struct {
	col       int
	row       int
	direction Direction
}

func (m *Marker) peekNext() (row int, col int) {
	switch m.direction {
	case North:
		return m.row - 1, m.col
	case East:
		return m.row, m.col + 1
	case South:
		return m.row + 1, m.col
	case West:
		return m.row, m.col - 1
	default:
		panic("unreachable")
	}
}

// ###

type Grid struct {
	lines   []string
	columns int
	rows    int
}

func newGrid(lines []string) Grid {
	return Grid{
		lines:   lines,
		columns: len(lines[0]),
		rows:    len(lines),
	}
}

func (g Grid) startMarker() Marker {
	for row := range g.rows {
		for col := range g.columns {
			val := string(g.lines[row][col])
			direction := parseDirection(val)
			if direction != None {
				return Marker{col, row, direction}
			}
		}
	}
	panic("marker not found")
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

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	grid := newGrid(lines)

	visited := determineVisited(grid)

	return len(visited)
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)

	initialGrid := newGrid(lines)
	initialMarker := initialGrid.startMarker()
	initialCoordinate := Coordinate{initialMarker.col, initialMarker.row}

	maxSteps := initialGrid.columns * initialGrid.rows

	guardPathCoordinates := determineVisited(initialGrid)
	loopsFound := 0

	for guardPathCoordinate := range guardPathCoordinates {
		if guardPathCoordinate == initialCoordinate {
			continue
		}

		stepCount := 0
		modifiedGrid := placeObstacleAt(initialGrid, guardPathCoordinate)
		hitDirections := NewSet[Direction]()
		marker := modifiedGrid.startMarker()

		for {

			nextRow, nextCol := marker.peekNext()
			char := modifiedGrid.charAt(nextRow, nextCol)

			if char == "" {
				break
			} else if stepCount >= maxSteps {
				loopsFound++
				break
			}

			if char == "#" {
				marker.direction = marker.direction.rotate()
			} else if char == "O" {
				if hitDirections.Contains(marker.direction) {
					loopsFound++
					break
				} else {
					hitDirections.Add(marker.direction)
					marker.direction = marker.direction.rotate()
				}
			} else {
				marker.col = nextCol
				marker.row = nextRow
				stepCount++
			}
		}

	}

	return loopsFound
}

func determineVisited(grid Grid) Set[Coordinate] {
	marker := grid.startMarker()

	visited := NewSet[Coordinate]()
	visited.Add(Coordinate{marker.col, marker.row})

	for {

		nextRow, nextCol := marker.peekNext()
		char := grid.charAt(nextRow, nextCol)

		if char == "" {
			break
		}

		if char == "#" {
			marker.direction = marker.direction.rotate()
		} else {
			marker.col = nextCol
			marker.row = nextRow
			visited.Add(Coordinate{marker.col, marker.row})
		}

	}
	return visited
}

func placeObstacleAt(g Grid, coordinate Coordinate) Grid {
	lineCopy := make([]string, len(g.lines))
	copy(lineCopy, g.lines)

	modifiedRow := []rune(lineCopy[coordinate.row])
	modifiedRow[coordinate.col] = 'O'
	lineCopy[coordinate.row] = string(modifiedRow)

	return newGrid(lineCopy)
}
