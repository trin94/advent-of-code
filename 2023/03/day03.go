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

// map with all marks "(x|y)"

func main() {
	file := "2023/03/input.txt"
	lines := readLinesFrom(file)
	numbers := parseNumbersFrom(lines)
	marks := parseMarksFrom(lines)

	part1Solution := solvePart1(numbers, marks)
	fmt.Printf("Part 1: %d\n", part1Solution)

	//part2Solution := solvePart2(games)
	//fmt.Printf("Part 2: %d\n", part2Solution)
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
			if unicode.IsDigit(character) {
				currentNumber += string(character)
				if currentNumberStartX == -1 {
					currentNumberStartX = x
					currentNumberEndX = x
				} else {
					currentNumberEndX++
				}
			}

			if !unicode.IsDigit(character) || (unicode.IsDigit(character) && x == len(line)-1) {
				if currentNumber != "" {
					value, _ := strconv.ParseInt(currentNumber, 10, 0)
					xStart := int32(currentNumberStartX)
					numbers = append(numbers, Number{
						xStart: xStart,
						xEnd:   max(xStart, int32(currentNumberEndX)),
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

func solvePart2() {

}
