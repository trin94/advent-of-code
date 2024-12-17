package p01

import (
	"slices"
	"strconv"
	"strings"
	"trin94/aoc/2024/inputs"
)

func solvePuzzle1(path string) (distances int) {
	lines := inputs.ReadLinesFrom(path)
	left, right := prepareSlices(lines)

	slices.Sort(left)
	slices.Sort(right)

	for idx := range left {
		l := left[idx]
		r := right[idx]

		if l < r {
			distances += r - l
		} else {
			distances += l - r
		}
	}

	return
}

func solvePuzzle2(path string) (score int) {
	lines := inputs.ReadLinesFrom(path)
	left, right := prepareSlices(lines)

	rightCounter := make(map[int]int, len(right))

	for _, r := range right {
		_, ok := rightCounter[r]
		if ok {
			rightCounter[r] += 1
		} else {
			rightCounter[r] = 1
		}
	}

	for _, l := range left {
		value, ok := rightCounter[l]
		if ok {
			score += l * value
		}
	}

	return
}

func prepareSlices(lines []string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range lines {
		split := strings.SplitN(line, "   ", 2)
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}
