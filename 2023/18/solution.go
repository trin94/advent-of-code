package main

import (
	"bytes"
	"strconv"
	"strings"
	"trin94/aoc/2023/inputs"
	"trin94/aoc/2023/utils"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Instruction struct {
	direction Direction
	amount    int
	color     string
}

func (i Instruction) UpdatedInstruction() Instruction {
	hex := i.color[1:6]
	directionStr := i.color[6:]

	var direction Direction
	if directionStr == "0" {
		direction = RIGHT
	} else if directionStr == "1" {
		direction = DOWN
	} else if directionStr == "2" {
		direction = LEFT
	} else if directionStr == "3" {
		direction = UP
	} else {
		panic("invalid direction" + directionStr)
	}

	amount, _ := strconv.ParseInt(hex, 16, 32)

	return Instruction{direction, int(amount), i.color}
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	instructions := parseInstructions(lines)
	gridLines := buildGrid(instructions)
	grid := utils.NewGrid(gridLines)

	totalArea := grid.Rows() * grid.Columns()
	areaPainted := 0

	start := utils.Coordinate{Col: 0, Row: 0}
	visited := utils.NewSet[utils.Coordinate]()

	visitNext := make([]utils.Coordinate, 0)
	visitNext = append(visitNext, start)

	adjacent := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for len(visitNext) > 0 {
		next := visitNext[0]
		visitNext = visitNext[1:]

		if visited.Contains(next) {
			continue
		}

		visited.Add(next)
		if grid.CharAtCoordinate(next) == '.' {
			areaPainted++
		}

		for _, adj := range adjacent {
			possibleNext := utils.Coordinate{Row: next.Row + adj[0], Col: next.Col + adj[1]}
			if visited.Contains(possibleNext) {
				continue
			}
			possibleChar := grid.CharAtCoordinate(possibleNext)
			if possibleChar == '#' || possibleChar == ' ' {
				continue
			}
			visitNext = append(visitNext, possibleNext)
		}
	}

	return totalArea - areaPainted
}

func parseInstructions(lines []string) []Instruction {
	result := make([]Instruction, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")

		directionChar := split[0]
		var direction Direction
		if directionChar == "L" {
			direction = LEFT
		} else if directionChar == "R" {
			direction = RIGHT
		} else if directionChar == "U" {
			direction = UP
		} else {
			direction = DOWN
		}

		amount, _ := strconv.Atoi(split[1])

		color := split[2]

		result[i] = Instruction{
			direction: direction,
			amount:    amount,
			color:     color[1 : len(color)-1],
		}
	}
	return result
}

func buildGrid(instructions []Instruction) []string {
	minRow, maxRow, minCol, maxCol := findOuterBounds(instructions)

	rows := maxRow - minRow + 3
	cols := maxCol - minCol + 3

	grid := make([][]byte, rows)
	for i := range grid {
		grid[i] = bytes.Repeat([]byte{'.'}, cols)
	}

	row, col := 1, 1
	for _, inst := range instructions {
		dr, dc := directionDelta(inst.direction)
		for i := 0; i < inst.amount; i++ {
			grid[row-minRow][col-minCol] = '#'
			row += dr
			col += dc
		}
	}
	grid[row-minRow][col-minCol] = '#'

	result := make([]string, rows)
	for i, r := range grid {
		result[i] = string(r)
	}

	return result
}

func findOuterBounds(instructions []Instruction) (int, int, int, int) {
	row, col := 0, 0

	minRow, maxRow := 0, 0
	minCol, maxCol := 0, 0

	for _, instruction := range instructions {
		switch instruction.direction {
		case UP:
			row -= instruction.amount
		case RIGHT:
			col += instruction.amount
		case DOWN:
			row += instruction.amount
		case LEFT:
			col -= instruction.amount
		}

		minRow = min(minRow, row)
		maxRow = max(maxRow, row)
		minCol = min(minCol, col)
		maxCol = max(maxCol, col)
	}

	return minRow, maxRow, minCol, maxCol
}

func directionDelta(dir Direction) (int, int) {
	switch dir {
	case UP:
		return -1, 0
	case DOWN:
		return 1, 0
	case RIGHT:
		return 0, 1
	case LEFT:
		return 0, -1
	}
	return 0, 0
}

func solvePuzzle2(path string) int {
	// lines := inputs.ReadLinesFrom(path)
	return 0
}
