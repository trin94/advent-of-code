package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file := "2023/08/input.txt"
	lines := readLinesFrom(file)

	instructionSet := lines[0]
	stepMapping := parseMappings(lines[2:])

	part1Solution := solvePart1(instructionSet, stepMapping)
	fmt.Printf("Part 1: %d\n", part1Solution)

	part2Solution := solvePart2(instructionSet, stepMapping)
	fmt.Printf("Part 2: %d\n", part2Solution)
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func parseMappings(lines []string) map[string]map[string]string {
	mapping := make(map[string]map[string]string)
	for _, line := range lines {
		fields := strings.Fields(line)
		k := fields[0]
		l := fields[2][1:4]
		r := fields[3][:3]

		value := make(map[string]string)
		value["L"] = l
		value["R"] = r

		mapping[k] = value
	}
	return mapping
}

func solvePart1(instructionSet string, mapping map[string]map[string]string) (stepCounter int) {
	var instructionIdx = 0
	var next = mapping["AAA"]
	for {
		lrStep := string(instructionSet[instructionIdx])
		target := next[lrStep]
		stepCounter++

		if target == "ZZZ" {
			break
		}

		next = mapping[target]

		if instructionIdx == len(instructionSet)-1 {
			instructionIdx = 0
		} else {
			instructionIdx++
		}
	}
	return
}

func solvePart2(instructionSet string, mapping map[string]map[string]string) (stepCounter int) {
	var loops []int
	for k := range mapping {
		if k[2] == 'A' {
			loops = append(loops, countStepsUntilLoop(instructionSet, mapping, k))
		}
	}
	return LCM(loops[0], loops[1], loops[2:]...) // thank u reddit
}

func countStepsUntilLoop(instructionSet string, mapping map[string]map[string]string, begin string) int {
	var stepCounter = 0
	var instructionIdx = 0
	var next = mapping[begin]

	var firstEnding = ""

	for {
		lrStep := string(instructionSet[instructionIdx])
		target := next[lrStep]
		stepCounter++

		if target[2] == 'Z' {
			if firstEnding == "" {
				firstEnding = target
			} else {
				return stepCounter / 2
			}
		}

		next = mapping[target]

		if instructionIdx == len(instructionSet)-1 {
			instructionIdx = 0
		} else {
			instructionIdx++
		}
	}
}

/*
https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
*/

// LCM find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// GCD greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
