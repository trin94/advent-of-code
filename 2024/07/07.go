package p07

import (
	"fmt"
	"strconv"
	"strings"
	"trin94/aoc/2024/inputs"
)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	matrix := buildOperatorMatrix()

	sum := 0

	for _, line := range lines {
		result, numbers := parseOperations(line)
		operatorCount := pow(2, len(numbers)-1)

		for operatorPossibility := 0; operatorPossibility < operatorCount; operatorPossibility++ {
			if accumulatesToUsingTwoOperators(result, numbers, matrix[operatorPossibility]) {
				sum += result
				break
			}
		}

	}

	return sum
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)

	sum := 0

	for _, line := range lines {
		result, numbers := parseOperations(line)
		for _, possibleCombination := range buildAccumulationSlices(numbers) {
			if accumulatesToUsingThreeOperators(result, possibleCombination) {
				sum += result
				break
			}
		}

	}

	return sum
}

func parseOperations(input string) (int, []int) {
	split := strings.SplitN(input, ": ", 2)
	result, _ := strconv.Atoi(split[0])

	remaining := strings.Split(split[1], " ")
	numbers := make([]int, 0)
	for _, value := range remaining {
		v, _ := strconv.Atoi(value)
		numbers = append(numbers, v)
	}

	return result, numbers
}

func buildOperatorMatrix() [][]bool {
	bitLength := 16
	twoToPowerOfBitLength := pow(2, 16)
	matrix := make([][]bool, twoToPowerOfBitLength)
	for nr := 0; nr < twoToPowerOfBitLength; nr++ {
		bits := make([]bool, bitLength)
		for bit := 0; bit < bitLength; bit++ {
			bits[bit] = (nr & (1 << bit)) > 0 // is bit set in nr
		}
		matrix[nr] = bits
	}
	return matrix
}

func accumulatesToUsingTwoOperators(target int, numbers []int, operators []bool) bool {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		nextNr := numbers[i]
		if operators[i-1] {
			result += nextNr
		} else {
			result *= nextNr
		}
		if result > target {
			return false
		}
	}
	return result == target
}

func buildAccumulationSlices(numbers []int) [][]int {
	operators := []int{0, 1, 2}

	numberCount := len(numbers)
	numberGaps := len(numbers) - 1

	rowLength := numberCount + numberGaps
	rowCount := pow(3, numberGaps)

	rows := make([][]int, rowCount)

	for i := 0; i < rowCount; i++ {
		rows[i] = make([]int, rowLength)
		operatorIndices := convertToBase(i, numberGaps)

		pos := 0
		for j := 0; j < numberCount; j++ {
			rows[i][pos] = numbers[j]
			pos++

			if j < numberGaps {
				rows[i][pos] = operators[operatorIndices[j]]
				pos++
			}
		}
	}

	return rows
}

func pow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func convertToBase(number int, length int) []int {
	base := 3
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[length-1-i] = number % base
		number /= base
	}
	return result
}

func accumulatesToUsingThreeOperators(target int, values []int) bool {
	current := values[0]

	length := len(values)
	idx := 1

	for {
		operator := values[idx]
		idx++
		next := values[idx]
		idx++

		if operator == 0 {
			current += next
		} else if operator == 1 {
			current *= next
		} else {
			current, _ = strconv.Atoi(fmt.Sprintf("%d%d", current, next))
		}

		if current > target {
			return false
		} else if idx >= length {
			break
		}
	}

	return current == target
}
