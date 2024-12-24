package main

import (
	"slices"
	"sort"
	"strings"
	"trin94/aoc/2024/inputs"
	"trin94/aoc/2024/utils"
)

type Connection struct {
	a, b string
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	connections := parseConnections(lines)
	groups := determineGroups(connections)

	groupsOfThree, result := make(map[string][]string), 0

	var isConnected = func(a, b string) bool {
		if network, found := groups[a]; found {
			return network.Contains(b)
		}
		return false
	}

	for computer, connected := range groups {
		if len(connected) < 2 {
			continue
		}

		computers := connected.Values()

		for i := 0; i < len(computers); i++ {
			for j := i + 1; j < len(computers); j++ {

				a := computers[i]
				b := computers[j]

				if isConnected(a, b) {
					subGroup := []string{computer, a, b}
					sort.Strings(subGroup)
					group := strings.Join(subGroup, "|")

					groupsOfThree[group] = subGroup
				}
			}
		}

	}

	for group := range groupsOfThree {
		if strings.HasPrefix(group, "t") || strings.Contains(group, "|t") {
			result++
		}
	}

	return result
}

func solvePuzzle2(path string) string {
	lines := inputs.ReadLinesFrom(path)
	connections := parseConnections(lines)
	groups := determineGroups(connections)

	networkSize, password := 0, ""

	var isConnected = func(a, b string) bool {
		if network, found := groups[a]; found {
			return network.Contains(b)
		}
		return false
	}

	for computer, network := range groups {
		if len(network) < 2 {
			continue
		}

		computers := network.Values()
		computers = append(computers, computer)
		slices.Sort(computers)

		largestGroup := findBiggestSubGroup(computers, isConnected)
		size := len(largestGroup)

		if size > networkSize {
			networkSize = size
			slices.Sort(largestGroup)
			password = strings.Join(largestGroup, ",")
		}
	}

	return password
}

func parseConnections(lines []string) []Connection {
	result := make([]Connection, len(lines))
	for i, line := range lines {
		computers := strings.Split(line, "-")
		result[i] = Connection{computers[0], computers[1]}
	}
	return result
}

func determineGroups(connections []Connection) map[string]utils.Set[string] {
	result := make(map[string]utils.Set[string])

	var connect = func(a, b string) {
		if network, found := result[a]; found {
			network.Add(b)
		} else {
			network := utils.NewSet[string]()
			network.Add(b)
			result[a] = network
		}
	}

	for _, connection := range connections {
		connect(connection.a, connection.b)
		connect(connection.b, connection.a)
	}

	return result
}

func findBiggestSubGroup(network []string, isConnected func(a, b string) bool) []string {
	for i := 0; i < len(network); i++ {
		for j := i + 1; j < len(network); j++ {
			if !isConnected(network[i], network[j]) {
				next := utils.Remove(utils.Clone(network), i)
				return findBiggestSubGroup(next, isConnected)
			}
		}
	}
	return network
}
