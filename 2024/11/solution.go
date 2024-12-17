package main

import (
	"strconv"
	"strings"
	"trin94/aoc/2024/inputs"
)

type StoneSteps struct {
	stone int
	steps int
}

var cache = make(map[StoneSteps]int)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	numbers := splitNumbers(lines[0])
	sum := 0
	for _, number := range numbers {
		sum += count(number, 25)
	}
	return sum
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	numbers := splitNumbers(lines[0])
	sum := 0
	for _, number := range numbers {
		sum += count(number, 75)
	}
	return sum
}

func splitNumbers(input string) []int {
	split := strings.Split(input, " ")
	result := make([]int, len(split))
	for i, s := range split {
		result[i], _ = strconv.Atoi(s)
	}
	return result
}

func count(stone int, steps int) int {
	k := StoneSteps{stone, steps}
	if cached, ok := cache[k]; ok {
		return cached
	}
	value := doCount(stone, steps)
	cache[k] = value
	return value
}

func doCount(stone int, steps int) int {
	if steps == 0 {
		return 1
	}
	if stone == 0 {
		return count(1, steps-1)
	}
	str := strconv.Itoa(stone)
	if len(str)%2 == 0 {
		idx := len(str) / 2
		first, _ := strconv.Atoi(str[:idx])
		second, _ := strconv.Atoi(str[idx:])
		return count(first, steps-1) + count(second, steps-1)
	}
	return count(stone*2024, steps-1)
}
