package main

import (
	"trin94/aoc/2023/inputs"
	"trin94/aoc/2023/utils"
)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	grid := utils.NewGrid(lines)
	moveStonesTowardsNorth(&grid)
	return calculateLoad(&grid)
}

func moveStonesTowardsNorth(grid *utils.Grid) {
	for col := 0; col < grid.Columns(); col++ {
		for row := 0; row < grid.Rows(); row++ {
			char := grid.CharAt(row, col)
			if char == 'O' {
				index := row
				for i := index - 1; i >= 0; i-- {
					above := grid.CharAt(i, col)
					if above == 'O' || above == '#' {
						break
					}
					index = i
				}
				temp := grid.CharAt(index, col)
				grid.SetCharAt('O', index, col)
				grid.SetCharAt(temp, row, col)
			}
		}
	}
}

func calculateLoad(grid *utils.Grid) (result int) {
	for row := 0; row < grid.Rows(); row++ {
		gravity := grid.Rows() - row
		for col := 0; col < grid.Columns(); col++ {
			if grid.CharAt(row, col) == 'O' {
				result += gravity
			}
		}
	}
	return
}

func solvePuzzle2(path string) (load int) {
	lines := inputs.ReadLinesFrom(path)
	grid := utils.NewGrid(lines)

	totalCycles := 1_000_000_000
	minCycleSize, maxCycleSize := 7, 15

	loads := make([]int, 0)

	for i := totalCycles; i > 0; i-- {

		for range 4 {
			moveStonesTowardsNorth(&grid)
			lines = utils.RotateClockwise(grid.Lines())
			grid = utils.NewGrid(lines)
		}

		load = calculateLoad(&grid)
		loads = append(loads, load)

		if len(loads) > minCycleSize {
			loadOccurredPreviouslyAtIndex := -1
			for j := len(loads) - 2; j >= 0 && j >= len(loads)-1-maxCycleSize; j-- {
				if loads[j] == load {
					loadOccurredPreviouslyAtIndex = j
					break
				}
			}
			if loadOccurredPreviouslyAtIndex != -1 {
				potentialCycleLength := len(loads) - 1 - loadOccurredPreviouslyAtIndex
				if potentialCycleLength >= minCycleSize {
					cycleFound := true
					for j := len(loads) - 1; j >= len(loads)-potentialCycleLength; j-- {
						if loads[j] != loads[j-potentialCycleLength] {
							cycleFound = false
							break
						}
					}
					if cycleFound {
						i %= potentialCycleLength
					}
				}
			}
		}

	}

	return
}
