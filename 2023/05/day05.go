package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type RuleSet struct {
	sourceRangeStart      int64
	destinationRangeStart int64
	rangeLength           int64
}

func (ruleSet RuleSet) mapMy(seed int64) int64 {
	minApplicableValue := ruleSet.sourceRangeStart
	maxApplicableValue := minApplicableValue + ruleSet.rangeLength - 1
	if seed >= minApplicableValue && seed <= maxApplicableValue {
		return ruleSet.destinationRangeStart - ruleSet.sourceRangeStart + seed
	}
	return -1
}

type Mapper struct {
	ruleSets []RuleSet
}

func (mapper Mapper) mapMy(seed int64) int64 {
	for _, ruleSet := range mapper.ruleSets {
		result := ruleSet.mapMy(seed)
		if result != -1 {
			return result
		}
	}
	return seed
}

func main() {
	file := "2023/05/input.txt"
	blocks := readBlocksFrom(file)

	seeds := parseSeeds(blocks[0])
	mapper := parseMapper(blocks[1:])

	part1Solution := solvePart1(seeds, mapper)
	fmt.Printf("Part 1: %d\n", part1Solution)

	part2Solution := solvePart2()
	fmt.Printf("Part 2: %d\n", part2Solution)
}

func readBlocksFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n\n")
}

func parseSeeds(line string) (seeds []int64) {
	line = strings.Split(line, ": ")[1]
	line = strings.TrimSpace(line)
	for _, seed := range strings.Split(line, " ") {
		value, _ := strconv.ParseInt(seed, 10, 0)
		seeds = append(seeds, value)
	}
	return
}

func parseMapper(blocks []string) (mappers []Mapper) {
	for _, block := range blocks {
		ruleBlock := strings.Split(block, ":\n")[1]

		var ruleSets []RuleSet
		for _, rule := range strings.Split(ruleBlock, "\n") {
			triple := strings.Split(rule, " ")
			destinationRangeStart, _ := strconv.ParseInt(triple[0], 10, 0)
			sourceRangeStart, _ := strconv.ParseInt(triple[1], 10, 0)
			rangeLength, _ := strconv.ParseInt(triple[2], 10, 0)

			ruleSets = append(ruleSets, RuleSet{
				sourceRangeStart:      sourceRangeStart,
				destinationRangeStart: destinationRangeStart,
				rangeLength:           rangeLength,
			})
		}
		mappers = append(mappers, Mapper{ruleSets: ruleSets})
	}

	return
}

func solvePart1(seeds []int64, mappers []Mapper) int64 {
	var result int64 = math.MaxInt64
	for _, seed := range seeds {
		value := seed
		for _, mapper := range mappers {
			value = mapper.mapMy(value)
		}
		result = min(result, value)
	}
	return result
}

func solvePart2() (result int) {
	return
}
