package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"trin94/aoc/2024/inputs"
	"trin94/aoc/2024/utils"
)

type Vertex struct {
	Key   utils.Coordinate
	Edges map[*Vertex]int
}

func solvePuzzle1(path string, size, corrupted int) int {
	lines := inputs.ReadLinesFrom(path)
	corruptedLocations := parseCorruptedLocations(lines)
	graph := buildGraph(size, toMap(corruptedLocations[:corrupted]))

	start := utils.Coordinate{Col: 0, Row: 0}
	end := utils.Coordinate{Col: size, Row: size}

	distances, err := Dijkstra(graph, start)
	if err != nil {
		panic(err)
	}

	return distances[end]
}

func solvePuzzle2(path string, size int) string {
	lines := inputs.ReadLinesFrom(path)
	corruptedLocations := parseCorruptedLocations(lines)

	start := utils.Coordinate{Col: 0, Row: 0}
	end := utils.Coordinate{Col: size, Row: size}

	cutOff := searchFirstCutOff(corruptedLocations, size, start, end)

	return fmt.Sprintf("%d,%d", cutOff.Col, cutOff.Row)
}

func parseCorruptedLocations(lines []string) []utils.Coordinate {
	result := make([]utils.Coordinate, len(lines))
	for i, line := range lines {
		xy := strings.Split(line, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		coordinate := utils.Coordinate{Col: x, Row: y}
		result[i] = coordinate
	}
	return result
}

func toMap(coordinates []utils.Coordinate) map[utils.Coordinate]struct{} {
	result := make(map[utils.Coordinate]struct{}, len(coordinates))
	for _, coordinate := range coordinates {
		result[coordinate] = struct{}{}
	}
	return result
}

func buildGraph(size int, corruptedLocations map[utils.Coordinate]struct{}) map[utils.Coordinate]*Vertex {
	vertices := make(map[utils.Coordinate]*Vertex, size)

	for r := 0; r <= size; r++ {
		for c := 0; c <= size; c++ {
			coordinate := utils.Coordinate{Row: r, Col: c}
			if _, corrupt := corruptedLocations[coordinate]; !corrupt {
				vertex := Vertex{Key: coordinate, Edges: make(map[*Vertex]int)}
				vertices[coordinate] = &vertex
			}
		}
	}

	adjacent := []utils.Coordinate{
		{Row: -1, Col: 0}, // north
		{Row: 0, Col: 1},  // east
		{Row: 1, Col: 0},  // south
		{Row: 0, Col: -1}, // south
	}

	for coordinate, vertex := range vertices {
		for _, adjacent := range adjacent {
			neighborCoordinate := utils.Coordinate{
				Col: coordinate.Col + adjacent.Col,
				Row: coordinate.Row + adjacent.Row,
			}
			if neighbor, found := vertices[neighborCoordinate]; found {
				vertex.Edges[neighbor] = 1
			}
		}
	}

	return vertices
}

// Dijkstra copied from https://reintech.io/blog/dijkstras-algorithm-in-go
func Dijkstra(graph map[utils.Coordinate]*Vertex, start utils.Coordinate) (map[utils.Coordinate]int, error) {
	_, ok := graph[start]
	if !ok {
		return nil, fmt.Errorf("start vertex %v not found", start)
	}

	distances := make(map[utils.Coordinate]int)
	for key := range graph {
		distances[key] = math.MaxInt32
	}
	distances[start] = 0

	var vertices []*Vertex
	for _, vertex := range graph {
		vertices = append(vertices, vertex)
	}

	for len(vertices) != 0 {
		sort.SliceStable(vertices, func(i, j int) bool {
			return distances[vertices[i].Key] < distances[vertices[j].Key]
		})

		vertex := vertices[0]
		vertices = vertices[1:]

		for adjacent, cost := range vertex.Edges {
			alt := distances[vertex.Key] + cost
			if alt < distances[adjacent.Key] {
				distances[adjacent.Key] = alt
			}
		}
	}

	return distances, nil
}

func searchFirstCutOff(corrupted []utils.Coordinate, size int, start, end utils.Coordinate) utils.Coordinate {

	for index, stepSize := 0, len(corrupted)/2; ; index = index + stepSize {

		coordinates := corrupted[:index+stepSize]
		graph := buildGraph(size, toMap(coordinates))

		distances, _ := Dijkstra(graph, start)
		distanceToEnd := distances[end]

		possibleToReachEnd := distanceToEnd != math.MaxInt32

		if stepSize == 1 && !possibleToReachEnd {
			return corrupted[index]
		}

		if !possibleToReachEnd {
			index -= stepSize
			stepSize /= 2
		}

	}
}
