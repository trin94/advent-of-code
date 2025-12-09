package main

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
	"trin94/aoc/2025/inputs"
	"trin94/aoc/2025/utils"
)

type BorderRange struct {
	min, max int
}

type BorderCollection struct {
	columnRanges map[int][]BorderRange
	rowRanges    map[int][]BorderRange
}

type Contracted struct {
	colSmallToBig, rowSmallToBig map[int]int
	coordinates                  []utils.Coordinate
	maxColumns, maxRows          int
}

func NewBorderCollection(coordinates []utils.Coordinate) BorderCollection {
	columnRanges := make(map[int][]BorderRange)
	rowRanges := make(map[int][]BorderRange)

	for i, length := 0, len(coordinates); i < length; i++ {
		first := coordinates[i]
		var second utils.Coordinate
		if i == length-1 {
			second = coordinates[0]
		} else {
			second = coordinates[i+1]
		}

		if first.Col == second.Col {
			ranges, exists := columnRanges[first.Col]
			if !exists {
				ranges = make([]BorderRange, 0)
			}
			ranges = append(ranges, BorderRange{min(first.Row, second.Row), max(first.Row, second.Row)})
			columnRanges[first.Col] = ranges
		}

		if first.Row == second.Row {
			ranges, exists := rowRanges[first.Row]
			if !exists {
				ranges = make([]BorderRange, 0)
			}
			ranges = append(ranges, BorderRange{min(first.Col, second.Col), max(first.Col, second.Col)})
			rowRanges[first.Row] = ranges
		}
	}

	return BorderCollection{
		columnRanges: columnRanges,
		rowRanges:    rowRanges,
	}
}

func (bc *BorderCollection) Contains(c utils.Coordinate) bool {
	if ranges, exists := bc.columnRanges[c.Col]; exists {
		for _, r := range ranges {
			if c.Row >= r.min && c.Row <= r.max {
				return true
			}
		}
	}

	if ranges, exists := bc.rowRanges[c.Row]; exists {
		for _, r := range ranges {
			if c.Col >= r.min && c.Col <= r.max {
				return true
			}
		}
	}

	return false
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	coordinates := parseCoordinates(lines)

	maxArea := 0

	for i := 0; i < len(coordinates); i++ {
		first := coordinates[i]
		for j := 0; j < len(coordinates); j++ {
			second := coordinates[j]
			current := areaBetween(first, second)
			maxArea = max(maxArea, current)
		}
	}

	return maxArea
}

func parseCoordinates(lines []string) []utils.Coordinate {
	result := make([]utils.Coordinate, len(lines))
	for idx, line := range lines {
		trimmed := strings.TrimSpace(line)
		split := strings.Split(trimmed, ",")
		first, _ := strconv.Atoi(split[0])
		second, _ := strconv.Atoi(split[1])
		result[idx] = utils.Coordinate{Col: first, Row: second}
	}
	return result
}

func areaBetween(c1, c2 utils.Coordinate) int {
	r := int(math.Abs(float64(c2.Row-c1.Row))) + 1
	c := int(math.Abs(float64(c2.Col-c1.Col))) + 1
	return r * c
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	coordinates := parseCoordinates(lines)
	contracted := contractCoordinates(coordinates)

	borders := NewBorderCollection(contracted.coordinates)
	outside := floodFill(contracted.maxRows, contracted.maxColumns, borders)

	maxArea := 0

	for i, length := 0, len(contracted.coordinates); i < length; i++ {
		first := contracted.coordinates[i]
		for j := i + 1; j < length; j++ {
			second := contracted.coordinates[j]
			if first.Row == second.Row && first.Col == second.Col {
				continue
			}

			if !isAreaLeakingOutside(first, second, outside) {
				origFirst := utils.Coordinate{
					Col: contracted.colSmallToBig[first.Col],
					Row: contracted.rowSmallToBig[first.Row],
				}
				origSecond := utils.Coordinate{
					Col: contracted.colSmallToBig[second.Col],
					Row: contracted.rowSmallToBig[second.Row],
				}
				maxArea = max(maxArea, areaBetween(origFirst, origSecond))
			}
		}
	}

	return maxArea
}

func contractCoordinates(coordinates []utils.Coordinate) Contracted {
	colSmallToBig := make(map[int]int)
	rowSmallToBig := make(map[int]int)

	colBigToSmall := make(map[int]int)
	rowBigToSmall := make(map[int]int)

	coordinatesClone := utils.Clone(coordinates)

	slices.SortFunc(coordinatesClone, func(a, b utils.Coordinate) int {
		return cmp.Or(
			cmp.Compare(a.Col, b.Col),
			cmp.Compare(a.Row, b.Row),
		)
	})

	counter := 1
	for _, coordinate := range coordinatesClone {
		big := coordinate.Col
		_, exists := colBigToSmall[big]
		if !exists {
			small := counter
			colBigToSmall[big] = small
			colSmallToBig[small] = big
			counter++
		}
	}
	maxColumns := counter

	slices.SortFunc(coordinatesClone, func(a, b utils.Coordinate) int {
		return cmp.Or(
			cmp.Compare(a.Row, b.Row),
			cmp.Compare(a.Col, b.Col),
		)
	})

	counter = 1
	for _, coordinate := range coordinatesClone {
		big := coordinate.Row
		_, exists := rowBigToSmall[big]
		if !exists {
			small := counter
			rowBigToSmall[big] = small
			rowSmallToBig[small] = big
			counter++
		}
	}
	maxRows := counter

	contracted := make([]utils.Coordinate, len(coordinates))
	for idx, coordinate := range coordinates {
		col := colBigToSmall[coordinate.Col]
		row := rowBigToSmall[coordinate.Row]
		contracted[idx] = utils.Coordinate{Col: col, Row: row}
	}

	return Contracted{
		colSmallToBig: colSmallToBig,
		rowSmallToBig: rowSmallToBig,
		coordinates:   contracted,
		maxColumns:    maxColumns,
		maxRows:       maxRows,
	}
}

func floodFill(maxRow, maxCol int, border BorderCollection) utils.Set[utils.Coordinate] {
	result := utils.NewSet[utils.Coordinate]()

	getAbove := func(c utils.Coordinate) (bool, utils.Coordinate) {
		return c.Row != 0, utils.Coordinate{Row: c.Row - 1, Col: c.Col}
	}

	getBelow := func(c utils.Coordinate) (bool, utils.Coordinate) {
		return c.Row != maxRow, utils.Coordinate{Row: c.Row + 1, Col: c.Col}
	}

	getLeft := func(c utils.Coordinate) (bool, utils.Coordinate) {
		return c.Col != 0, utils.Coordinate{Row: c.Row, Col: c.Col - 1}
	}

	getRight := func(c utils.Coordinate) (bool, utils.Coordinate) {
		return c.Col != maxCol, utils.Coordinate{Row: c.Row, Col: c.Col + 1}
	}

	visited := utils.NewSet[utils.Coordinate]()

	checkNext := make([]utils.Coordinate, 0)
	checkNext = append(checkNext, utils.Coordinate{Row: 0, Col: 0})

	nextGetters := []func(utils.Coordinate) (bool, utils.Coordinate){getAbove, getBelow, getLeft, getRight}

	for len(checkNext) > 0 {
		next := checkNext[0]
		checkNext = checkNext[1:]

		for _, function := range nextGetters {
			valid, adjacent := function(next)
			if !valid || visited.Contains(adjacent) {
				continue
			}

			visited.Add(adjacent)

			if !border.Contains(adjacent) {
				result.Add(adjacent)
				checkNext = append(checkNext, adjacent)
			}
		}
	}

	return result
}

func isAreaLeakingOutside(c1, c2 utils.Coordinate, outside utils.Set[utils.Coordinate]) bool {
	minRow := min(c1.Row, c2.Row)
	maxRow := max(c1.Row, c2.Row)
	minCol := min(c1.Col, c2.Col)
	maxCol := max(c1.Col, c2.Col)

	for col := minCol; col <= maxCol; col++ {
		if outside.Contains(utils.Coordinate{Row: minRow, Col: col}) {
			return true
		}
		if outside.Contains(utils.Coordinate{Row: maxRow, Col: col}) {
			return true
		}
	}

	for row := minRow + 1; row < maxRow; row++ {
		if outside.Contains(utils.Coordinate{Row: row, Col: minCol}) {
			return true
		}
		if outside.Contains(utils.Coordinate{Row: row, Col: maxCol}) {
			return true
		}
	}

	return false
}
