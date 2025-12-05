package main

import (
	"slices"
	"strconv"
	"strings"
	"trin94/aoc/2025/inputs"
)

type Range struct {
	begin, end int
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	ranges, ingredients := parseInput(lines)

	fresh := 0

	for _, ingredient := range ingredients {
		if isInRange(ingredient, ranges) {
			fresh++
		}
	}

	return fresh
}

func parseInput(lines []string) ([]Range, []int) {
	ranges := make([]Range, 0)
	ingredients := make([]int, 0)

	readingIngredients := false

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			readingIngredients = true
			continue
		}
		if readingIngredients {
			ingredient, _ := strconv.Atoi(trimmedLine)
			ingredients = append(ingredients, ingredient)
		} else {
			split := strings.Split(trimmedLine, "-")
			begin, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			ranges = append(ranges, Range{begin, end})
		}
	}

	return ranges, ingredients
}

func isInRange(number int, ranges []Range) bool {
	for _, r := range ranges {
		if number >= r.begin && number <= r.end {
			return true
		}
	}
	return false
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	ranges, _ := parseInput(lines)

	slices.SortFunc(ranges, func(a, b Range) int {
		if a.begin == b.begin {
			return 0
		}
		if a.begin <= b.begin {
			return -1
		}
		return 1
	})

	sum := 0

	for _, r := range untangleRanges(ranges) {
		sum += r.end - r.begin + 1
	}

	return sum
}

func untangleRanges(input []Range) []Range {
	result := make([]Range, 0)

	for i := 0; i < len(input); i++ {
		rangeA := input[i]
		begin := rangeA.begin
		end := rangeA.end

		for j := i + 1; j < len(input); j++ {
			rangeB := input[j]
			if end < rangeB.begin {
				break
			}
			end = max(end, rangeB.end)
			i++
		}

		result = append(result, Range{begin, end})
	}

	return result
}
