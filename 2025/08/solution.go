package main

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
	"trin94/aoc/2025/inputs"
	"trin94/aoc/2025/utils"
)

type Coordinate struct {
	x, y, z int
}

func NewCoordinate(x, y, z int) Coordinate {
	return Coordinate{x, y, z}
}

func (c *Coordinate) EuclideanDistance(other Coordinate) float64 {
	deltaX := other.x - c.x
	deltaY := other.y - c.y
	deltaZ := other.z - c.z
	powers := deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ
	result := math.Sqrt(float64(powers))
	return roundToDecimal(result, 2)
}

func roundToDecimal(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

type Connection struct {
	first, second Coordinate
	distance      float64
}

func solvePuzzle1(path string, connectionCount int) int {
	lines := inputs.ReadLinesFrom(path)
	coordinates := parseCoordinates(lines)
	distances := calculateDistances(coordinates)

	slices.SortFunc(distances, func(a, b Connection) int {
		return cmp.Compare(a.distance, b.distance)
	})

	circuits := make(map[Coordinate]*[]Coordinate)
	for _, coordinate := range coordinates {
		slice := make([]Coordinate, 0)
		circuits[coordinate] = &slice
	}

	circuits = buildCircuits(coordinates, distances[0:connectionCount])

	deduplicatedCircuits := utils.NewSet[*[]Coordinate]()
	for _, circuit := range circuits {
		deduplicatedCircuits.Add(circuit)
	}

	deduplicatedCircuitsSorted := deduplicatedCircuits.Values()
	slices.SortFunc(deduplicatedCircuitsSorted, func(a, b *[]Coordinate) int {
		return cmp.Compare(len(*b), len(*a))
	})

	first := len(*deduplicatedCircuitsSorted[0])
	second := len(*deduplicatedCircuitsSorted[1])
	third := len(*deduplicatedCircuitsSorted[2])

	return first * second * third
}

func parseCoordinates(lines []string) []Coordinate {
	result := make([]Coordinate, len(lines))
	for idx, line := range lines {
		trimmed := strings.TrimSpace(line)
		split := strings.Split(trimmed, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		result[idx] = Coordinate{x, y, z}
	}
	return result
}

func calculateDistances(coordinates []Coordinate) []Connection {
	result := make([]Connection, 0)
	for i := 0; i < len(coordinates); i++ {
		first := coordinates[i]
		for j := i + 1; j < len(coordinates); j++ {
			second := coordinates[j]
			distance := first.EuclideanDistance(second)
			connection := Connection{first, second, distance}
			result = append(result, connection)
		}
	}
	return result
}

func buildCircuits(coordinates []Coordinate, connections []Connection) map[Coordinate]*[]Coordinate {
	circuits := make(map[Coordinate]*[]Coordinate)
	for _, coordinate := range coordinates {
		slice := []Coordinate{coordinate}
		circuits[coordinate] = &slice
	}

	for _, connection := range connections {
		first := circuits[connection.first]
		second := circuits[connection.second]

		if first == second {
			continue
		}

		*first = append(*first, *second...)

		for _, coordinate := range *second {
			circuits[coordinate] = first
		}
	}

	return circuits
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	coordinates := parseCoordinates(lines)
	distances := calculateDistances(coordinates)

	slices.SortFunc(distances, func(a, b Connection) int {
		return cmp.Compare(a.distance, b.distance)
	})

	circuits := make(map[Coordinate]*[]Coordinate)
	for _, coordinate := range coordinates {
		slice := make([]Coordinate, 0)
		circuits[coordinate] = &slice
	}

	first, second := buildCircuitsUntilOneCircuit(coordinates, distances)
	return first * second
}

func buildCircuitsUntilOneCircuit(coordinates []Coordinate, connections []Connection) (int, int) {
	circuits := make(map[Coordinate]*[]Coordinate)
	for _, coordinate := range coordinates {
		slice := []Coordinate{coordinate}
		circuits[coordinate] = &slice
	}

	counter := len(circuits)

	for _, connection := range connections {
		first := circuits[connection.first]
		second := circuits[connection.second]

		if first == second {
			continue
		}

		*first = append(*first, *second...)

		for _, coordinate := range *second {
			circuits[coordinate] = first
		}

		counter--

		if counter == 1 {
			return connection.first.x, connection.second.x
		}
	}

	return -1, -1
}
