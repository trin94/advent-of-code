package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	inp01sample := "resources/day01.txt"
	lines := ReadLines(inp01sample)

	part1Solution := part1(lines)
	fmt.Printf("Part 1: %d\n", part1Solution)

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
