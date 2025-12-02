package main

import (
	"strconv"
	"strings"
	"trin94/aoc/2025/inputs"
)

type Range struct {
	begin, end int
}

func newRange(input string) Range {
	numRangeSplit := strings.Split(strings.TrimSpace(input), "-")
	begin, _ := strconv.ParseInt(numRangeSplit[0], 10, 0)
	end, _ := strconv.ParseInt(numRangeSplit[1], 10, 0)
	return Range{int(begin), int(end)}
}

func (r Range) invalidIds() []int {
	result := make([]int, 0)
	for i := r.begin; i <= r.end; i++ {
		number := strconv.Itoa(i)
		if len(number)%2 != 0 {
			continue
		}
		first := number[len(number)/2:]
		last := number[:len(number)/2]
		if first == last {
			result = append(result, i)
		}
	}
	return result
}

func (r Range) invalidIdsExtended() []int {
	result := make([]int, 0)
	for i := r.begin; i <= r.end; i++ {
		number := strconv.Itoa(i)
		length := len(number)
		for j := 1; j <= length; j++ {
			chunks := chunkString(number, j)
			first := chunks[0]
			allEqual := allMatch(chunks, func(s string) bool {
				return first == s
			})
			if len(chunks) > 1 && allEqual {
				result = append(result, i)
				break
			}
		}
	}
	return result
}

func chunkString(s string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

func allMatch[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	line := lines[0]

	sum := 0
	for _, r := range parseRanges(line) {
		for _, invalid := range r.invalidIds() {
			sum += invalid
		}
	}
	return sum
}

func parseRanges(line string) []Range {
	ranges := make([]Range, 0)
	for numRange := range strings.SplitSeq(line, ",") {
		r := newRange(numRange)
		ranges = append(ranges, r)
	}
	return ranges
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	line := lines[0]

	sum := 0
	for _, r := range parseRanges(line) {
		for _, invalid := range r.invalidIdsExtended() {
			sum += invalid
		}
	}
	return sum
}
