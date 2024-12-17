package p02

import (
	"strconv"
	"strings"
	"trin94/aoc/2024/inputs"
)

func solvePuzzle1(path string) (distances int) {
	lines := inputs.ReadLinesFrom(path)
	counter := 0

	for _, line := range lines {
		numbers := splitAsIntSlice(line)

		if verifyLine(numbers) {
			counter++
		}
	}

	return int(counter)
}

func solvePuzzle2(path string) (score int) {
	lines := inputs.ReadLinesFrom(path)
	counter := 0

	for _, line := range lines {
		numbers := splitAsIntSlice(line)

		idx := 0
		for idx < len(numbers) {
			cloned := clone(numbers)
			adapted := remove(cloned, idx)

			idx++

			if verifyLine(adapted) {
				counter++
				break
			}
		}
	}

	return int(counter)
}

func verifyLine(numbers []int) bool {
	first := numbers[0]
	second := numbers[1]

	if first == second {
		return false
	}

	ascending := first < second

	idx := 0
	length := len(numbers) - 1

	for idx < length {
		first = numbers[idx]
		second = numbers[idx+1]

		if ascending && !(second > first && first+3 >= second) {
			return false
		}

		if !ascending && !(second < first && first-3 <= second) {
			return false
		}

		idx++
	}

	return true
}

func splitAsIntSlice(slice string) []int {
	result := make([]int, 0)
	for _, value := range strings.Split(slice, " ") {
		nr, _ := strconv.Atoi(value)
		result = append(result, nr)
	}
	return result
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func clone(slice []int) []int {
	into := make([]int, len(slice))
	copy(into, slice)
	return into
}
