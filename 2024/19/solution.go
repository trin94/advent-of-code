package main

import (
	"strings"
	"trin94/aoc/2024/inputs"
)

func solvePuzzle1(path string) (result int) {
	lines := inputs.ReadLinesFrom(path)
	patterns, designs := parsePuzzleInput(lines)
	for _, design := range designs {
		if canProduce(design, &patterns) {
			result++
		}
	}
	return
}

func solvePuzzle2(path string) (result int) {
	lines := inputs.ReadLinesFrom(path)
	patterns, designs := parsePuzzleInput(lines)
	cache := make(map[string]int)
	for _, design := range designs {
		result += canProduceCount(design, &patterns, &cache)
	}
	return
}

func parsePuzzleInput(lines []string) (patterns, designs []string) {
	patternLine := lines[0]
	for _, pattern := range strings.Split(patternLine, ", ") {
		patterns = append(patterns, pattern)
	}
	designs = lines[2:]
	return patterns, designs
}

func canProduce(design string, patterns *[]string) bool {
	if len(design) == 0 {
		return true
	}
	for _, pattern := range *patterns {
		trimmed := strings.TrimSuffix(design, pattern)
		if trimmed == design {
			continue
		}
		if canProduce(trimmed, patterns) {
			return true
		}
	}
	return false
}

func canProduceCount(design string, patterns *[]string, cache *map[string]int) int {
	if len(design) == 0 {
		return 1
	}
	if possibilities, found := (*cache)[design]; found {
		return possibilities
	}
	possibilities := 0
	for _, pattern := range *patterns {
		trimmed := strings.TrimSuffix(design, pattern)
		if trimmed == design {
			continue
		}
		possibilities += canProduceCount(trimmed, patterns, cache)
	}
	if possibilities > 0 {
		(*cache)[design] = possibilities
	}
	return possibilities
}
