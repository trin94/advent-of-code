package main

import (
	"cmp"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"trin94/aoc/2025/inputs"
)

type MathOperation int

const (
	Add MathOperation = iota
	Multiply
)

type MathProblem struct {
	numbers   []int
	operation MathOperation
}

func NewMathProblem(items []string) MathProblem {
	numbers := make([]int, len(items)-1)
	for i := 0; i < len(numbers); i++ {
		nr, _ := strconv.Atoi(items[i])
		numbers[i] = nr
	}
	var operation MathOperation
	if items[len(items)-1] == "*" {
		operation = Multiply
	} else {
		operation = Add
	}
	return MathProblem{
		numbers:   numbers,
		operation: operation,
	}
}

func NewCephalopodMathProblem(items []string) MathProblem {
	numbers := make([]int, 0)
	rowCount := len(items) - 1
	colCount := len(items[0])

	for col := 0; col < colCount; col++ {
		concat := ""
		for row := 0; row < rowCount; row++ {
			concat += string(items[row][col])
		}
		nr, _ := strconv.Atoi(strings.TrimSpace(concat))
		numbers = append(numbers, nr)
	}

	var operation MathOperation
	if strings.TrimSpace(items[len(items)-1]) == "*" {
		operation = Multiply
	} else {
		operation = Add
	}
	return MathProblem{
		numbers:   numbers,
		operation: operation,
	}
}

var WHITESPACE = regexp.MustCompile("\\s+")

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	problems := parseProblems(lines)
	return solve(problems)
}

func solve(problems []MathProblem) int {
	sum := 0
	for _, problem := range problems {
		current := 0
		if problem.operation == Add {
			for _, n := range problem.numbers {
				current += n
			}
		} else {
			current = 1
			for _, n := range problem.numbers {
				current *= n
			}
		}
		sum += current
	}
	return sum
}

func parseProblems(lines []string) []MathProblem {
	split := splitLines(lines)
	rowCount := len(split)
	colCount := len(split[0])

	result := make([]MathProblem, colCount)
	for c := 0; c < colCount; c++ {
		items := make([]string, rowCount)
		for r := 0; r < rowCount; r++ {
			items[r] = split[r][c]
		}
		result[c] = NewMathProblem(items)
	}

	return result
}

func splitLines(lines []string) [][]string {
	result := make([][]string, len(lines))
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		split := WHITESPACE.Split(trimmed, -1)
		result[i] = split
	}
	return result
}

func solvePuzzle2(path string) int {
	lines := readLinesCustom(path)
	columnIndexes := parseColumnBreakIndexes(lines)
	blocks := parseBlocks(lines, columnIndexes)
	problems := parseCephalopodMathProblems(blocks)
	return solve(problems)
}

func readLinesCustom(path string) []string {
	inputByteStream, err := os.ReadFile(path)
	if err != nil {
		panic("Could not read file from " + path)
	}
	content := string(inputByteStream)
	content = strings.TrimSuffix(content, "\n")
	return strings.Split(content, "\n")
}

func parseColumnBreakIndexes(lines []string) []int {
	result := make([]int, 0)
	line := lines[0]
	for c, char := range line {
		if char == ' ' {
			isColumnBreak := true
			for _, otherLine := range lines[1:] {
				otherChar := otherLine[c]
				if otherChar != ' ' {
					isColumnBreak = false
					break
				}
			}
			if isColumnBreak {
				result = append(result, c)
			}
		}
	}
	return result
}

func parseBlocks(lines []string, columnBreakIndexes []int) [][]string {
	blocks := make([][]string, 0)
	block := make([]string, 0)

	rowCount := len(lines)
	index := 0

	for _, breakIndex := range columnBreakIndexes {
		for row := 0; row < rowCount; row++ {
			a := lines[row][index:breakIndex]
			block = append(block, a)
		}
		index = breakIndex + 1
		blocks = append(blocks, block)
		block = make([]string, 0)
	}

	longestLine := slices.MaxFunc(lines, func(a, b string) int {
		i := len(a)
		i2 := len(b)
		return cmp.Compare(i, i2)
	})
	longest := len(longestLine)

	for row := 0; row < rowCount; row++ {
		slice := lines[row]
		a := (slice + "        ")[index:longest]
		block = append(block, a)
	}

	return append(blocks, block)
}

func parseCephalopodMathProblems(blocks [][]string) []MathProblem {
	result := make([]MathProblem, 0)
	for _, block := range blocks {
		problem := NewCephalopodMathProblem(block)
		result = append(result, problem)
	}
	return result
}
