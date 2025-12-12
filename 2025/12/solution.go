package main

import (
	"strconv"
	"strings"
	"trin94/aoc/2025/inputs"
	"trin94/aoc/2025/utils"
)

type Shape struct {
	index int
	grid  utils.Grid
}

func NewShape(lines []string) Shape {
	index, _ := strconv.Atoi(string(lines[0][0]))
	grid := utils.NewGrid(lines[1:])
	return Shape{index, grid}
}

type Region struct {
	width, length int
	indexList     []int
}

func NewRegion(line string) Region {
	split := strings.Split(line, ":")
	measurements := strings.Split(split[0], "x")
	width, _ := strconv.Atoi(measurements[0])
	length, _ := strconv.Atoi(measurements[1])

	shapeIndexes := strings.Split(strings.TrimSpace(split[1]), " ")
	indexList := make([]int, len(shapeIndexes))
	for i, value := range shapeIndexes {
		indexList[i], _ = strconv.Atoi(value)
	}
	return Region{width, length, indexList}
}

func (r *Region) IndexCount() int {
	return len(r.indexList)
}

func (r *Region) GetIndex(idx int) int {
	if idx < 0 || idx >= r.IndexCount() {
		panic("index out of range")
	}
	return r.indexList[idx]
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	_, regions := parseInput(lines)

	sum := 0
	for _, region := range regions {
		count := 0
		for i := 0; i < region.IndexCount(); i++ {
			count += region.GetIndex(i)
		}
		shapeArea := count * (3 * 3)
		availableArea := region.width * region.length
		if shapeArea <= availableArea {
			sum++
		}
	}
	return sum
}

func parseInput(lines []string) ([]Shape, []Region) {
	sections := make([][]string, 0)
	section := make([]string, 0)
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			sections = append(sections, section)
			section = make([]string, 0)
			continue
		}
		section = append(section, line)
	}

	shapes := make([]Shape, 0)
	for _, line := range sections {
		shapes = append(shapes, NewShape(line))
	}

	regions := make([]Region, 0)
	for _, line := range section {
		regions = append(regions, NewRegion(line))
	}

	return shapes, regions
}
