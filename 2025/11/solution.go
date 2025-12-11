package main

import (
	"strings"
	"trin94/aoc/2025/inputs"
)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	graph := parseGraph(lines)
	return solve(graph)
}

func parseGraph(lines []string) map[string][]string {
	graph := make(map[string][]string)
	for _, line := range lines {
		split := strings.Split(line, ":")
		source := strings.TrimSpace(split[0])
		targets := make([]string, 0)
		for _, target := range strings.Split(strings.TrimSpace(split[1]), " ") {
			targets = append(targets, strings.TrimSpace(target))
		}
		graph[source] = targets
	}
	return graph
}

func solve(graph map[string][]string) int {
	memo := make(map[string]int)

	var count func(node string) int
	count = func(node string) int {
		if node == "out" {
			return 1
		}
		if cached, ok := memo[node]; ok {
			return cached
		}
		result := 0
		for _, target := range graph[node] {
			result += count(target)
		}
		memo[node] = result
		return result
	}

	return count("you")
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	graph := parseGraph(lines)
	return solve2(graph)
}

type MemoKey struct {
	node    string
	fftSeen bool
	dacSeen bool
}

func solve2(graph map[string][]string) int {
	memo := make(map[MemoKey]int)

	var count func(node string, fftSeen, dacSeen bool) int
	count = func(node string, fftSeen, dacSeen bool) int {
		if node == "fft" {
			fftSeen = true
		} else if node == "dac" {
			dacSeen = true
		}
		if node == "out" {
			if dacSeen && fftSeen {
				return 1
			}
			return 0
		}
		key := MemoKey{node, fftSeen, dacSeen}
		if cached, ok := memo[key]; ok {
			return cached
		}
		total := 0
		for _, next := range graph[node] {
			total += count(next, fftSeen, dacSeen)
		}
		memo[key] = total
		return total
	}

	return count("svr", false, false)
}
