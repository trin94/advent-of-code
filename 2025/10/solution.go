package main

import (
	"trin94/aoc/2025/inputs"
)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	machines := parseMachines(lines)

	sum := 0

	for _, machine := range machines {
		sum += machine.CalculateFewestButtonPressesToTurnOn()
	}

	return sum
}

func parseMachines(lines []string) []Machine {
	machines := make([]Machine, len(lines))
	for i, line := range lines {
		machines[i] = NewMachine(line)
	}
	return machines
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	machines := parseMachines(lines)

	sum := 0

	for _, machine := range machines {
		sum += machine.CalculateFewestButtonPressesToReachVoltage()
	}

	return sum
}
