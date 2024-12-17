package main

import (
	"math"
	"trin94/aoc/2024/inputs"
	"trin94/aoc/2024/utils"
)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	universe := utils.NewGrid(lines)
	galaxies := findGalaxies(universe)
	ySpace, xSpace := detectNotExpandingSpace(universe)

	sum := 0

	for g1 := 0; g1 < len(galaxies); g1++ {
		for g2 := g1 + 1; g2 < len(galaxies); g2++ {
			sum += distanceBetween(galaxies[g1], galaxies[g2], ySpace, xSpace, 1)
		}
	}

	return sum
}

func solvePuzzle2(path string, expandFactor int) int {
	lines := inputs.ReadLinesFrom(path)
	universe := utils.NewGrid(lines)
	galaxies := findGalaxies(universe)
	ySpace, xSpace := detectNotExpandingSpace(universe)

	sum := 0

	for g1 := 0; g1 < len(galaxies); g1++ {
		for g2 := g1 + 1; g2 < len(galaxies); g2++ {
			sum += distanceBetween(galaxies[g1], galaxies[g2], ySpace, xSpace, expandFactor)
		}
	}

	return sum
}

func findGalaxies(grid utils.Grid) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)
	for r := range grid.Rows() {
		for c := range grid.Columns() {
			if grid.CharAt(r, c) == '#' {
				result = append(result, utils.NewCoordinate(c, r))
			}
		}
	}
	return result
}

func detectNotExpandingSpace(grid utils.Grid) (map[int]struct{}, map[int]struct{}) {
	ySpace := make(map[int]struct{})
	xSpace := make(map[int]struct{})

	for r := range grid.Rows() {
		for c := range grid.Columns() {
			if grid.CharAt(r, c) != '.' {
				ySpace[r] = struct{}{}
				xSpace[c] = struct{}{}
			}
		}
	}

	return ySpace, xSpace
}

func distanceBetween(c1, c2 utils.Coordinate, notExpandingY, notExpandingX map[int]struct{}, factor int) int {
	rDelta := int(math.Abs(float64(c2.Row - c1.Row)))
	cDelta := int(math.Abs(float64(c2.Col - c1.Col)))

	spaceMultiplier := 0
	rMin := min(c1.Row, c2.Row)
	for i := rMin; i < rMin+rDelta; i++ {
		if _, found := notExpandingY[i]; !found {
			spaceMultiplier++
		}
	}

	cMin := min(c1.Col, c2.Col)
	for i := cMin; i < cMin+cDelta; i++ {
		if _, found := notExpandingX[i]; !found {
			spaceMultiplier++
		}
	}

	return rDelta + cDelta + spaceMultiplier*max(1, factor-1)
}
