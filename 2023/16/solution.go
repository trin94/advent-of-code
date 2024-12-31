package main

import (
	"sync"
	"trin94/aoc/2023/inputs"
	"trin94/aoc/2023/utils"
)

type Tick struct {
	coordinate utils.Coordinate
	direction  utils.Direction
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	grid := utils.NewGrid(lines)

	initial := Tick{utils.NewCoordinate(0, 0), utils.East}

	return calculateEnergizedCount(grid, initial)
}

func calculateEnergizedCount(grid utils.Grid, initial Tick) int {
	visited, visitedCoordinates := utils.NewSet[Tick](), utils.NewSet[utils.Coordinate]()

	next := make([]Tick, 0)
	next = append(next, initial)

	appendNorthOf := func(t Tick) { next = append(next, Tick{t.coordinate.North(), utils.North}) }
	appendSouthOf := func(t Tick) { next = append(next, Tick{t.coordinate.South(), utils.South}) }
	appendWestOf := func(t Tick) { next = append(next, Tick{t.coordinate.West(), utils.West}) }
	appendEastOf := func(t Tick) { next = append(next, Tick{t.coordinate.East(), utils.East}) }

	for len(next) > 0 {
		current := next[0]
		next = next[1:]

		currentChar := grid.CharAtCoordinate(current.coordinate)
		if currentChar == ' ' || visited.Contains(current) {
			continue
		}

		visited.Add(current)
		visitedCoordinates.Add(current.coordinate)

		switch current.direction {
		case utils.North:
			switch currentChar {
			case '.':
				appendNorthOf(current)
			case '/':
				appendEastOf(current)
			case '\\':
				appendWestOf(current)
			case '-':
				appendWestOf(current)
				appendEastOf(current)
			case '|':
				appendNorthOf(current)
			}
		case utils.East:
			switch currentChar {
			case '.':
				appendEastOf(current)
			case '/':
				appendNorthOf(current)
			case '\\':
				appendSouthOf(current)
			case '-':
				appendEastOf(current)
			case '|':
				appendNorthOf(current)
				appendSouthOf(current)
			}
		case utils.South:
			switch currentChar {
			case '.':
				appendSouthOf(current)
			case '/':
				appendWestOf(current)
			case '\\':
				appendEastOf(current)
			case '-':
				appendWestOf(current)
				appendEastOf(current)
			case '|':
				appendSouthOf(current)
			}
		case utils.West:
			switch currentChar {
			case '.':
				appendWestOf(current)
			case '/':
				appendSouthOf(current)
			case '\\':
				appendNorthOf(current)
			case '-':
				appendWestOf(current)
			case '|':
				appendSouthOf(current)
				appendNorthOf(current)
			}
		}
	}

	return len(visitedCoordinates)
}

func solvePuzzle2(path string) (result int) {
	lines := inputs.ReadLinesFrom(path)
	grid := utils.NewGrid(lines)

	cols, rows := grid.Columns(), grid.Rows()

	initials := make([]Tick, 0)
	for c := range cols {
		initials = append(initials, Tick{utils.Coordinate{Col: c}, utils.South})
		initials = append(initials, Tick{utils.Coordinate{Col: c, Row: rows - 1}, utils.North})
	}
	for r := range rows {
		initials = append(initials, Tick{utils.Coordinate{Row: r}, utils.East})
		initials = append(initials, Tick{utils.Coordinate{Row: r, Col: cols - 1}, utils.West})
	}

	results := make(chan int, len(initials))

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(initials))

	for _, initial := range initials {
		go func(grid utils.Grid, initial Tick) {
			defer waitGroup.Done()
			results <- calculateEnergizedCount(grid, initial)
		}(grid, initial)
	}

	waitGroup.Wait()
	close(results)

	for r := range results {
		if r > result {
			result = r
		}
	}

	return
}
