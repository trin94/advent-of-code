package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	xStart int32
	xEnd   int32
	y      int32
	value  int32
}

func (number Number) isConnectedTo(mark Mark) bool {
	xInRange := mark.x >= number.xStart-1 && mark.x <= number.xEnd+1
	yInRange := mark.y == number.y-1 || mark.y == number.y || mark.y == number.y+1
	return xInRange && yInRange
}

type Mark struct {
	x         int32
	y         int32
	character rune
}

func main() {
	file := "2023/03/input.txt"
	lines := readLinesFrom(file)
	numbers := parseNumbersFrom(lines)
	marks := parseMarksFrom(lines)

	part1Solution := solvePart1(numbers, marks)
	fmt.Printf("Part 1: %d\n", part1Solution)

	part2Solution := solvePart2(numbers, marks)
	fmt.Printf("Part 2: %d\n", part2Solution)
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func parseNumbersFrom(lines []string) []Number {
	var numbers []Number

	for y, line := range lines {
		var currentNumber string
		var currentNumberStartX = -1
		var currentNumberEndX = -1

		for x, character := range line {
			isDigit := unicode.IsDigit(character)

			if isDigit {
				currentNumber += string(character)
				if currentNumberStartX == -1 {
					currentNumberStartX = x
					currentNumberEndX = x
				} else {
					currentNumberEndX++
				}
			}

			isLastDigitInLine := isDigit && x == len(line)-1
			if isLastDigitInLine || !isDigit {
				if currentNumber != "" {
					value, _ := strconv.ParseInt(currentNumber, 10, 0)
					numbers = append(numbers, Number{
						xStart: int32(currentNumberStartX),
						xEnd:   int32(currentNumberEndX),
						y:      int32(y),
						value:  int32(value),
					})
				}

				currentNumber = ""
				currentNumberStartX = -1
				currentNumberEndX = -1
			}
		}
	}

	return numbers
}

func parseMarksFrom(lines []string) []Mark {
	var marks []Mark
	for y, line := range lines {
		for x, character := range line {
			if !unicode.IsDigit(character) && character != '.' {
				marks = append(marks, Mark{
					x:         int32(x),
					y:         int32(y),
					character: character,
				})
			}
		}
	}
	return marks
}

func solvePart1(numbers []Number, marks []Mark) int32 {
	var sum int32
	for _, number := range numbers {
		for _, mark := range marks {
			if number.isConnectedTo(mark) {
				sum += number.value
				continue
			}
		}
	}
	return sum
}

func solvePart2(numbers []Number, marks []Mark) int32 {
	starMarks := filter(marks, func(mark Mark) bool { return mark.character == '*' })

	var sum int32

	for _, mark := range starMarks {
		var factors []int32
		for _, number := range numbers {
			if number.isConnectedTo(mark) {
				factors = append(factors, number.value)
			}
		}
		if len(factors) == 2 {
			sum += factors[0] * factors[1]
		}
	}

	return sum
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
