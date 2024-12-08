package p02

import (
	"os"
	"strconv"
	"strings"
)

func solvePuzzle1(path string) (distances int64) {
	lines := readLinesFrom(path)
	counter := 0

	for _, line := range lines {
		numbers := splitAsIntSlice(line)

		if verifyLine(numbers) {
			counter++
		}
	}

	return int64(counter)
}

func solvePuzzle2(path string) (score int64) {
	lines := readLinesFrom(path)
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

	return int64(counter)
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func verifyLine(numbers []int64) bool {
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

func splitAsIntSlice(slice string) []int64 {
	result := make([]int64, 0)
	for _, value := range strings.Split(slice, " ") {
		nr, _ := strconv.ParseInt(value, 10, 0)
		result = append(result, nr)
	}
	return result
}

func remove(slice []int64, s int) []int64 {
	return append(slice[:s], slice[s+1:]...)
}

func clone(slice []int64) []int64 {
	into := make([]int64, len(slice))
	copy(into, slice)
	return into
}
