package main

import (
	"slices"
	"trin94/aoc/2025/inputs"
	"trin94/aoc/2025/utils"
)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	grid := utils.NewGrid(lines)
	start := grid.FindCoordinateWithChar('S')
	return countSplits(start, grid)
}

func countSplits(start utils.Coordinate, grid utils.Grid) int {
	visited := utils.NewSet[utils.Coordinate]()
	visitNext := make([]utils.Coordinate, 0)
	visitNext = append(visitNext, start)

	addToVisitNext := func(c utils.Coordinate) {
		reachedBottom := grid.CharAtCoordinate(c) == ' '
		if reachedBottom {
			return
		}
		if !visited.Contains(c) {
			visitNext = append(visitNext, c)
		}
	}

	for len(visitNext) > 0 {
		current := visitNext[0]
		visitNext = visitNext[1:]

		nextChar := grid.CharAtCoordinate(current)

		if nextChar == '.' || nextChar == 'S' {
			dropped := utils.Coordinate{Col: current.Col, Row: current.Row + 1}
			addToVisitNext(dropped)
			continue
		}

		if nextChar == '^' {
			visited.Add(current)
			left := utils.Coordinate{Col: current.Col - 1, Row: current.Row}
			right := utils.Coordinate{Col: current.Col + 1, Row: current.Row}
			addToVisitNext(left)
			addToVisitNext(right)
		}

	}

	return len(visited)
}

func solvePuzzle2(path string) int {
	// Basic idea: amount of timelines = total split count. First, we find all coordinates where a split occurs.
	// Then, we sort them so that we start processing from the bottom. We use a cache to only calculate splits per '^'
	// coordinate once. In the end, our cache contains the amount of splits that would happen if we continue downwards
	// for each split coordinate. We add 1 for the 'current' timeline and we're done.

	lines := inputs.ReadLinesFrom(path)
	grid := utils.NewGrid(lines)
	start := grid.FindCoordinateWithChar('S')

	splits := make(map[utils.Coordinate]int)
	splitCoordinates := findSplitCoordinates(&grid)
	slices.SortFunc(splitCoordinates, func(a, b utils.Coordinate) int {
		return b.Row - a.Row
	})

	for _, split := range splitCoordinates {
		splits[split] = countSplitsWithCache(split, grid, splits)
	}

	return splits[start] + 1
}

func findSplitCoordinates(grid *utils.Grid) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)
	for row := 0; row < grid.Rows; row++ {
		for col := 0; col < grid.Columns; col++ {
			character := grid.CharAt(row, col)
			if character == '^' || character == 'S' {
				result = append(result, utils.Coordinate{Row: row, Col: col})
			}
		}
	}
	return result
}

func countSplitsWithCache(start utils.Coordinate, grid utils.Grid, cache map[utils.Coordinate]int) int {
	visited := utils.NewSet[utils.Coordinate]()
	visitNext := make([]utils.Coordinate, 0)
	visitNext = append(visitNext, start)

	addToVisitNext := func(c utils.Coordinate) {
		reachedBottom := grid.CharAtCoordinate(c) == ' '
		if reachedBottom {
			return
		}
		if !visited.Contains(c) {
			visitNext = append(visitNext, c)
		}
	}

	cacheSum := 0

	for len(visitNext) > 0 {
		current := visitNext[0]
		visitNext = visitNext[1:]

		nextChar := grid.CharAtCoordinate(current)

		if nextChar == '.' || nextChar == 'S' {
			dropped := utils.Coordinate{Col: current.Col, Row: current.Row + 1}
			addToVisitNext(dropped)
			continue
		}

		if nextChar == '^' {
			if splitsForCurrent, cacheHit := cache[current]; cacheHit {
				cacheSum += splitsForCurrent
				continue
			}

			visited.Add(current)
			left := utils.Coordinate{Col: current.Col - 1, Row: current.Row}
			right := utils.Coordinate{Col: current.Col + 1, Row: current.Row}
			addToVisitNext(left)
			addToVisitNext(right)
		}

	}

	return len(visited) + cacheSum
}
