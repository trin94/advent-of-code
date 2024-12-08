package p09

import (
	"os"
	"strconv"
	"strings"
)

func solvePuzzle1(path string) int {
	lines := readLinesFrom(path)

	sum := 0

	for _, line := range lines {
		numbers := mapToNumbers(line)

		subCalculations := make([][]int, 0)
		subCalculations = append(subCalculations, numbers)

		values := numbers
		for {
			differences, allZeros := calculateDifferences(values)
			subCalculations = append(subCalculations, differences)
			if allZeros {
				break
			}
			values = differences
		}

		change := 0
		for i := len(subCalculations) - 1; i > 0; i-- {
			calculation := subCalculations[i]
			lastValue := calculation[len(calculation)-1]
			change += lastValue
		}

		sum += numbers[len(numbers)-1] + change
	}

	return sum
}

func solvePuzzle2(path string) int {
	lines := readLinesFrom(path)

	sum := 0

	for _, line := range lines {
		numbers := mapToNumbers(line)

		subCalculations := make([][]int, 0)
		subCalculations = append(subCalculations, numbers)

		values := numbers
		for {
			differences, allZeros := calculateDifferences(values)
			subCalculations = append(subCalculations, differences)
			if allZeros {
				break
			}
			values = differences
		}

		change := 0
		for i := len(subCalculations) - 1; i >= 0; i-- {
			calculation := subCalculations[i]
			firstValue := calculation[0]
			change = firstValue - change
		}

		sum += change
	}

	return sum
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func mapToNumbers(input string) []int {
	split := strings.Split(input, " ")
	result := make([]int, len(split))
	for i, str := range split {
		result[i], _ = strconv.Atoi(str)
	}
	return result
}

func calculateDifferences(numbers []int) ([]int, bool) {
	allZeros := true
	result := make([]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		current := numbers[i]
		next := numbers[i+1]
		difference := next - current
		result[i] = difference
		if difference != 0 {
			allZeros = false
		}
	}
	return result, allZeros
}
