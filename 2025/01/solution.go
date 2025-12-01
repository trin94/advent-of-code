package main

import (
	"strconv"
	"trin94/aoc/2025/inputs"
)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)

	position := 50
	password := 0

	for _, line := range lines {
		operation := line[0]
		amount, _ := strconv.ParseInt(line[1:], 10, 0)

		direction := 1
		if operation == 'L' {
			direction = -1
		}

		position = (position + (direction * int(amount)) + 100) % 100

		if position == 0 {
			password++
		}
	}

	return password
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)

	position := 50
	password := 0

	for _, line := range lines {
		operation := line[0]
		amount, _ := strconv.ParseInt(line[1:], 10, 0)

		direction := 1
		if operation == 'L' {
			direction = -1
		}

		for i := 0; i < int(amount); i++ {
			position = (position + direction + 100) % 100

			if position == 0 {
				password++
			}
		}
	}

	return password
}
