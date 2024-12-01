package p01

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func solvePuzzle1(path string) (distances int64) {
	lines := readLinesFrom(path)
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

func solvePuzzle2(path string) (score int64) {
	lines := readLinesFrom(path)
	left, right := prepareSlices(lines)

	rightCounter := make(map[int64]int64, len(right))

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

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func prepareSlices(lines []string) ([]int64, []int64) {
	var left []int64
	var right []int64

	for _, line := range lines {
		split := strings.SplitN(line, "   ", 2)
		l, _ := strconv.ParseInt(split[0], 10, 0)
		r, _ := strconv.ParseInt(split[1], 10, 0)
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}
