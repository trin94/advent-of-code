package main

import (
	"fmt"
	"strconv"
	"trin94/aoc/2024/inputs"
	"unicode"
)

type Code struct {
	value       string
	numericPart int
}

var cache = make(map[string]int)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	codes := readCodes(lines)

	sum, depth := 0, 2

	for _, code := range codes {
		initial := calculateFullPath(code.value)
		length := evolve(initial, depth)
		numValue := code.numericPart
		sum += numValue * length
	}

	return sum

}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	codes := readCodes(lines)

	sum, depth := 0, 25

	for _, code := range codes {
		initial := calculateFullPath(code.value)
		length := evolve(initial, depth)
		numValue := code.numericPart
		sum += numValue * length
	}

	return sum
}

func readCodes(lines []string) []Code {
	codes := make([]Code, len(lines))
	for i, line := range lines {
		numberConcat := make([]rune, 0)
		for _, r := range line {
			if unicode.IsDigit(r) {
				numberConcat = append(numberConcat, r)
			}
		}
		numericPart, _ := strconv.Atoi(string(numberConcat))
		codes[i] = Code{value: line, numericPart: numericPart}
	}
	return codes
}

func evolve(sequence string, depth int) int {
	if depth == 0 {
		return len(sequence)
	}

	cacheKey := fmt.Sprintf("%s|%d", sequence, depth)

	if val, found := cache[cacheKey]; found {
		return val
	}

	length, current := 0, 'A'

	for _, next := range sequence {
		nextSequence := calculatePath(current, next)
		length += evolve(nextSequence, depth-1)
		current = next
	}

	cache[cacheKey] = length
	return length
}

func calculateFullPath(value string) string {
	current, result := 'A', ""
	for _, next := range value {
		result += calculatePath(current, next)
		current = next
	}
	return result
}

func calculatePath(before, after rune) string {
	switch {
	case before == after:
		return "A"
	default:
		return pathMap[before][after] + "A"
	}
}

var pathMap = map[rune]map[rune]string{
	'A': {
		'0': "<",
		'1': "^<<",
		'2': "<^",
		'3': "^",
		'4': "^^<<",
		'5': "<^^",
		'6': "^^",
		'7': "^^^<<",
		'8': "<^^^",
		'9': "^^^",
		'^': "<",
		'>': "v",
		'<': "v<<",
		'v': "<v",
	},
	'0': {
		'A': ">",
		'1': "^<",
		'2': "^",
		'3': "^>",
		'4': "^^<",
		'5': "^^",
		'6': "^^>",
		'7': "^^^<",
		'8': "^^^",
		'9': "^^^>",
	},
	'1': {
		'A': ">>v",
		'0': ">v",
		'2': ">",
		'3': ">>",
		'4': "^",
		'5': "^>",
		'6': "^>>",
		'7': "^^",
		'8': "^^>",
		'9': "^^>>",
	},
	'2': {
		'A': "v>",
		'0': "v",
		'1': "<",
		'3': ">",
		'4': "<^",
		'5': "^",
		'6': "^>",
		'7': "<^^",
		'8': "^^",
		'9': "^^>",
	},
	'3': {
		'A': "v",
		'0': "<v",
		'1': "<<",
		'2': "<",
		'4': "<<^",
		'5': "<^",
		'6': "^",
		'7': "<<^^",
		'8': "<^^",
		'9': "^^",
	},
	'4': {
		'A': ">>vv",
		'0': ">vv",
		'1': "v",
		'2': "v>",
		'3': "v>>",
		'5': ">",
		'6': ">>",
		'7': "^",
		'8': "^>",
		'9': "^>>",
	},
	'5': {
		'A': "vv>",
		'0': "vv",
		'1': "<v",
		'2': "v",
		'3': "v>",
		'4': "<",
		'6': ">",
		'7': "<^",
		'8': "^",
		'9': "^>",
	},
	'6': {
		'A': "vv",
		'0': "<vv",
		'1': "<<v",
		'2': "<v",
		'3': "v",
		'4': "<<",
		'5': "<",
		'7': "<<^",
		'8': "<^",
		'9': "^",
	},
	'7': {
		'A': ">>vvv",
		'0': ">vvv",
		'1': "vv",
		'2': "vv>",
		'3': "vv>>",
		'4': "v",
		'5': "v>",
		'6': "v>>",
		'8': ">",
		'9': ">>",
	},
	'8': {
		'A': "vvv>",
		'0': "vvv",
		'1': "<vv",
		'2': "vv",
		'3': "vv>",
		'4': "<v",
		'5': "v",
		'6': "v>",
		'7': "<",
		'9': ">",
	},
	'9': {
		'A': "vvv",
		'0': "<vvv",
		'1': "<<vv",
		'2': "<vv",
		'3': "vv",
		'4': "<<v",
		'5': "<v",
		'6': "v",
		'7': "<<",
		'8': "<",
	},
	'<': {
		'A': ">>^",
		'v': ">",
		'^': ">^",
		'>': ">>",
	},
	'>': {
		'A': "^",
		'v': "<",
		'^': "<^",
		'<': "<<",
	},
	'^': {
		'A': ">",
		'v': "v",
		'>': "v>",
		'<': "v<",
	},
	'v': {
		'A': "^>",
		'^': "^",
		'>': ">",
		'<': "<",
	},
}
