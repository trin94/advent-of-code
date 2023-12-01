package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	inp01sample := "resources/day01.txt"
	lines := ReadLines(inp01sample)

	part1Solution := part1(lines)
	fmt.Printf("Part 1: %d\n", part1Solution)

	part2Solution := part2(lines)
	fmt.Printf("Part 2: %d\n", part2Solution)

}

func part1(lines []string) int64 {
	var sum int64 = 0

	for _, line := range lines {
		var firstDigit string
		var lastDigit string

		for _, character := range line {
			if !unicode.IsDigit(character) {
				continue
			}
			if firstDigit == "" {
				firstDigit = string(character)
			}
			lastDigit = string(character)
		}

		valuePerLine, _ := strconv.ParseInt(firstDigit+lastDigit, 10, 0)
		sum += valuePerLine
	}

	return sum
}

func part2(lines []string) int64 {

	var manipulatedLine []string
	for _, line := range lines {
		line = strings.ReplaceAll(line, "one", "O1E")
		line = strings.ReplaceAll(line, "One", "O1E")
		line = strings.ReplaceAll(line, "onE", "O1E")

		line = strings.ReplaceAll(line, "two", "T2O")
		line = strings.ReplaceAll(line, "Two", "T2O")
		line = strings.ReplaceAll(line, "twO", "T2O")

		line = strings.ReplaceAll(line, "three", "T3E")
		line = strings.ReplaceAll(line, "Three", "T3E")
		line = strings.ReplaceAll(line, "threE", "T3E")

		line = strings.ReplaceAll(line, "four", "F4R")
		line = strings.ReplaceAll(line, "Four", "F4R")
		line = strings.ReplaceAll(line, "fouR", "F4R")

		line = strings.ReplaceAll(line, "five", "F5E")
		line = strings.ReplaceAll(line, "Five", "F5E")
		line = strings.ReplaceAll(line, "fivE", "F5E")

		line = strings.ReplaceAll(line, "six", "S6X")
		line = strings.ReplaceAll(line, "Six", "S6X")
		line = strings.ReplaceAll(line, "siX", "S6X")

		line = strings.ReplaceAll(line, "seven", "S7N")
		line = strings.ReplaceAll(line, "Seven", "S7N")
		line = strings.ReplaceAll(line, "seveN", "S7N")

		line = strings.ReplaceAll(line, "eight", "E8T")
		line = strings.ReplaceAll(line, "Eight", "E8T")
		line = strings.ReplaceAll(line, "eighT", "E8T")

		line = strings.ReplaceAll(line, "nine", "N9E")
		line = strings.ReplaceAll(line, "Nine", "N9E")
		line = strings.ReplaceAll(line, "ninE", "N9E")

		manipulatedLine = append(manipulatedLine, line)
	}

	return part1(manipulatedLine)
}
