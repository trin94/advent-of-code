package main

import (
	"slices"
	"strings"
	"trin94/aoc/2023/inputs"
	"trin94/aoc/2023/utils"
)

func solvePuzzle1(path string) (result int) {
	lines := inputs.ReadLinesFrom(path)
	patterns := parsePatterns(lines)

	for _, pattern := range patterns {
		if index := findMirrorIndex(pattern); index >= 0 {
			result += 100 * index
		}
		if index := findMirrorIndex(rotate(pattern)); index >= 0 {
			result += index
		}
	}

	return result
}

func solvePuzzle2(path string) (result int) {
	lines := inputs.ReadLinesFrom(path)
	patterns := parsePatterns(lines)

	for _, pattern := range patterns {
		index := findMirrorIndex(pattern)
		for _, fixed := range fixSmudges(pattern) {
			newIndex := findMirrorIndexButIgnore(fixed, []int{index})
			if newIndex > 0 && newIndex != index {
				result += 100 * newIndex
				break
			}
		}
		patternRotated := rotate(pattern)
		index = findMirrorIndex(patternRotated)
		for _, fixed := range fixSmudges(patternRotated) {
			newIndex := findMirrorIndexButIgnore(fixed, []int{index})
			if newIndex > 0 && newIndex != index {
				result += newIndex
				break
			}
		}
	}

	return result
}

func parsePatterns(lines []string) [][]string {
	chunk, patterns := make([]string, 0), make([][]string, 0)
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, chunk)
			chunk = make([]string, 0)
		} else {
			chunk = append(chunk, line)
		}
	}
	return append(patterns, chunk)
}

// findMirrorIndex returns the index of a reflection or -1 if none exists
func findMirrorIndex(lines []string) int {
	return findMirrorIndexButIgnore(lines, []int{})
}

func findMirrorIndexButIgnore(lines []string, ignoreIndexes []int) int {
	previousLine, indexes := "", make([]int, 0)
	for index, line := range lines {
		if line == previousLine {
			indexes = append(indexes, index)
		}
		previousLine = line
	}
	indexes = utils.Filter(indexes, func(index int) bool {
		return !slices.Contains(ignoreIndexes, index)
	})
	for _, index := range indexes {
		iToBegin, iToEnd, isPerfectReflection := index-1, index, true
		for begin, end := 0, len(lines); iToBegin >= begin && iToEnd < end; iToBegin, iToEnd = iToBegin-1, iToEnd+1 {
			if lines[iToBegin] != lines[iToEnd] {
				isPerfectReflection = false
				break
			}
		}
		if isPerfectReflection {
			return index
		}
	}
	return -1
}

func rotate(lines []string) []string {
	maxRows, maxColumns := len(lines), len(lines[0])
	result := make([]string, maxColumns)
	for col := 0; col < maxColumns; col++ {
		var sb strings.Builder
		for row := maxRows - 1; row >= 0; row-- {
			sb.WriteByte(lines[row][col])
		}
		result[col] = sb.String()
	}
	return result
}

func fixSmudges(lines []string) (result [][]string) {
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if 1 == editDistance(lines[i], lines[j]) {
				clone := utils.Clone(lines)
				clone[i] = clone[j]
				if findMirrorIndex(clone) >= 0 {
					result = append(result, clone)
				}
			}
		}
	}
	return result
}

func editDistance(a, b string) (result int) {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			result++
		}
	}
	return result
}
