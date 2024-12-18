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

const MaxDistance = math.MaxInt32

type Vertex struct {
	Key   utils.Coordinate
	Edges map[utils.Coordinate]int
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
				vertex := Vertex{Key: coordinate, Edges: make(map[utils.Coordinate]int)}
				vertices[coordinate] = &vertex
			}
		}
	}

	adjacent := []utils.Coordinate{
		{Row: -1, Col: 0}, // north
		{Row: 0, Col: 1},  // east
		{Row: 1, Col: 0},  // south
		{Row: 0, Col: -1}, // west
	}

	for coordinate, vertex := range vertices {
		for _, adjacent := range adjacent {
			neighborCoordinate := utils.Coordinate{
				Col: coordinate.Col + adjacent.Col,
				Row: coordinate.Row + adjacent.Row,
			}
			if _, found := vertices[neighborCoordinate]; found {
				vertex.Edges[neighborCoordinate] = 1
			}
		}
	}

	return vertices
}

// Dijkstra copied and adapted from https://reintech.io/blog/dijkstras-algorithm-in-go
func Dijkstra(graph map[utils.Coordinate]*Vertex, start utils.Coordinate) (map[utils.Coordinate]int, error) {
	_, ok := graph[start]
	if !ok {
		return nil, fmt.Errorf("start vertex %v not found", start)
	}

	distances := make(map[utils.Coordinate]int)
	for key := range graph {
		distances[key] = MaxDistance
	}
	distances[start] = 0

	var vertices []*Vertex
	for _, vertex := range graph {
		vertices = append(vertices, vertex)
	}

	for len(vertices) != 0 {
		sort.Slice(vertices, func(i, j int) bool {
			return distances[vertices[i].Key] < distances[vertices[j].Key]
		})

		vertex := vertices[0]
		vertices = vertices[1:]

		for adjacent, cost := range vertex.Edges {
			alt := distances[vertex.Key] + cost
			if alt < distances[adjacent] {
				distances[adjacent] = alt
			}
		}
	}

	return distances, nil
}

func searchFirstCutOff(corrupted []utils.Coordinate, size int, start, end utils.Coordinate) utils.Coordinate {

	center := (len(corrupted) - 1) / 2
	window, index := center, center

	for window > 1 {

		graph := buildGraph(size, toMap(corrupted[0:index]))
		distances, _ := Dijkstra(graph, start)
		distanceToEnd := distances[end]
		possibleToReachEnd := distanceToEnd != MaxDistance

		window = window / 2

		if possibleToReachEnd {
			index += window
		} else {
			index -= window
		}

	}

	return corrupted[index]

}
