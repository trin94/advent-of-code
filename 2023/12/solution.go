package main

import (
	"fmt"
	"strconv"
	"strings"
	"trin94/aoc/2023/inputs"
)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)

	cache := make(map[string]int)

	sum := 0

	for _, conditionRecord := range lines {
		condition, groups := parseRecord(conditionRecord)
		sum += count(condition, groups, cache)
	}

	return sum
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)

	cache := make(map[string]int)

	sum := 0
	for _, conditionRecord := range lines {
		condition, groups := parseRecord(conditionRecord)
		condition, groups = expand(condition, groups)
		sum += count(condition, groups, cache)
	}

	return sum
}

func parseRecord(line string) (string, []int) {
	fields := strings.Fields(line)
	lineStr := fields[0]
	numbers := strings.Split(fields[1], ",")
	groups := make([]int, len(numbers))
	for i, number := range numbers {
		groups[i], _ = strconv.Atoi(number)
	}

	return lineStr, groups
}

func expand(line string, groups []int) (outString string, outGroup []int) {
	outString = strings.Join([]string{line, line, line, line, line}, "?")
	outGroup = make([]int, len(groups)*5)
	for i, group := range groups {
		for j := 0; j < len(groups)*5; j = j + len(groups) {
			outGroup[i+j] = group
		}
	}
	return
}

// count created based on https://www.youtube.com/watch?v=g3Ms5e7Jdqo, thanks dude :)
func count(source string, groups []int, cache map[string]int) int {
	if source == "" {
		// if empty string and no groups left
		if len(groups) == 0 {
			return 1
		}
		// if empty string but groups left
		return 0
	}

	// if no groups but string left
	if len(groups) == 0 {
		sharpCount := strings.Count(source, "#")
		if sharpCount == 0 {
			return 1
		}
		return 0
	}

	cacheKey := fmt.Sprintf("%s|%v", source, groups)

	if cachedValue, found := cache[cacheKey]; found {
		return cachedValue
	}

	nextChar := source[0]

	result := 0

	if nextChar == '.' || nextChar == '?' {
		result += count(source[1:], groups, cache)
	}

	if nextChar == '#' || nextChar == '?' {

		charCountLeft := len(source)
		nextGroupSize := groups[0]

		if nextGroupSize <= charCountLeft && // enough springs left?
			0 == strings.Count(source[:nextGroupSize], ".") && // must not contain operational springs until next group
			(nextGroupSize == charCountLeft || source[nextGroupSize] != '#') { // either no springs left or no damaged springs immediately after

			if nextGroupSize == charCountLeft {
				result += count("", groups[1:], cache)
			} else {
				result += count(source[nextGroupSize+1:], groups[1:], cache)
			}

		}
	}

	cache[cacheKey] = result

	return result
}
