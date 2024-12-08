package p10

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Coordinate struct {
	row, col int
}

// ###

type Tile int

const (
	Start Tile = iota
	Vertical
	Horizontal
	NorthEast
	NorthWest
	SouthWest
	SouthEast
	None
)

// ###

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

func (g Grid) TileAt(row int, col int) Tile {
	if row >= g.rows || row < 0 {
		return None
	}
	if col >= g.columns || col < 0 {
		return None
	}
	switch g.lines[row][col] {
	case 'S':
		return Start
	case '|':
		return Vertical
	case '-':
		return Horizontal
	case 'L':
		return NorthEast
	case 'J':
		return NorthWest
	case '7':
		return SouthWest
	case 'F':
		return SouthEast
	default:
		return None
	}
}

func (g Grid) Start() (row, col int) {
	for row := range g.rows {
		for col := range g.columns {
			if g.TileAt(row, col) == Start {
				return col, row
			}
		}
	}
	panic("Could not find starting position")
}

func (g Grid) detectTileAt(row int, col int) Tile {
	north := g.TileAt(row-1, col)
	south := g.TileAt(row+1, col)
	west := g.TileAt(row, col-1)
	east := g.TileAt(row, col+1)

	connectedNorth := north == Vertical || north == SouthWest || north == SouthEast
	connectedSouth := south == Vertical || south == NorthWest || south == NorthEast
	connectedWest := west == Horizontal || west == NorthEast || west == SouthEast
	connectedEast := east == Horizontal || east == NorthWest || east == SouthWest

	if connectedNorth && connectedSouth {
		return Vertical
	} else if connectedWest && connectedEast {
		return Horizontal
	} else if connectedNorth && connectedEast {
		return NorthEast
	} else if connectedEast && connectedSouth {
		return SouthEast
	} else if connectedSouth && connectedWest {
		return SouthWest
	} else if connectedWest && connectedNorth {
		return NorthWest
	} else {
		panic(fmt.Sprintf("Cannot detect tile at col=%d, row=%d, ", col, row))
	}
}

func (g Grid) Directions(row int, col int) (opt1Row, opt1Col, opt2Row, opt2Col int) {
	tile := g.TileAt(row, col)
	if tile == Start {
		tile = g.detectTileAt(row, col)
	}

	switch tile {
	case Vertical:
		opt1Row = row - 1
		opt1Col = col
		opt2Row = row + 1
		opt2Col = col
	case Horizontal:
		opt1Row = row
		opt1Col = col + 1
		opt2Row = row
		opt2Col = col - 1
	case NorthEast:
		opt1Row = row - 1
		opt1Col = col
		opt2Row = row
		opt2Col = col + 1
	case NorthWest:
		opt1Row = row - 1
		opt1Col = col
		opt2Row = row
		opt2Col = col - 1
	case SouthWest:
		opt1Row = row + 1
		opt1Col = col
		opt2Row = row
		opt2Col = col - 1
	case SouthEast:
		opt1Row = row
		opt1Col = col + 1
		opt2Row = row + 1
		opt2Col = col
	default:
		panic("Out of bounds")
	}
	return
}

func solvePuzzle1(path string) int {
	lines := readLinesFrom(path)
	grid := NewGrid(lines)
	visitedCoordinates := findLoop(grid)
	return len(visitedCoordinates) / 2
}

func solvePuzzle2(path string) int {
	lines := readLinesFrom(path)
	grid := NewGrid(lines)
	visitedCoordinates := findLoop(grid)

	// Source: https://www.reddit.com/r/adventofcode/comments/18evyu9/comment/kcso138/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
	// thank you! :)

	// https://en.wikipedia.org/wiki/Shoelace_formula
	// https://en.wikipedia.org/wiki/Pick%27s_theorem

	// Calculate the entire area with the shoelace formula
	// then use picks theorem to extract just the inner area
	// (rearrange formula so that we're looking for i)

	sum := 0

	for i := 0; i < len(visitedCoordinates); i++ {
		n1 := visitedCoordinates[i]
		n2 := visitedCoordinates[(i+1)%len(visitedCoordinates)]

		col1, row1 := n1.col, n1.row
		col2, row2 := n2.col, n2.row

		sum += col1*row2 - row1*col2
	}

	area := int(math.Abs(float64(sum / 2)))

	return area - len(visitedCoordinates)/2 + 1
}

func findLoop(grid Grid) []Coordinate {
	rowStart, colStart := grid.Start()

	visited := make([]Coordinate, 0)
	visited = append(visited, Coordinate{rowStart, colStart})

	nextCol := colStart
	nextRow := rowStart

	prevCol := colStart
	prevRow := rowStart

	for {
		opt1Row, opt1Col, opt2Row, opt2Col := grid.Directions(nextRow, nextCol)

		opt1Visited := opt1Col == prevCol && opt1Row == prevRow

		prevCol = nextCol
		prevRow = nextRow

		if opt1Visited {
			nextCol = opt2Col
			nextRow = opt2Row
		} else {
			nextCol = opt1Col
			nextRow = opt1Row
		}

		visited = append(visited, Coordinate{nextRow, nextCol})

		if nextCol == colStart && nextRow == rowStart {
			break
		}
	}

	return visited
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}
