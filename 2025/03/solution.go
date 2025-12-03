package main

import (
	"strconv"
	"strings"
	"trin94/aoc/2025/inputs"
)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	sum := 0

	for _, battery := range lines {
		sum += calculateJoltage(battery, 2)
	}

	return sum
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	sum := 0

	for _, line := range lines {
		sum += calculateJoltage(line, 12)
	}

	return sum
}

func calculateJoltage(battery string, keepCells int) int {
	removeCount := len(battery) - keepCells
	batterySplit := strings.Split(battery, "")
	optimizedEnergy := remove(batterySplit, removeCount)
	finalEnergyStr := strings.Join(optimizedEnergy, "")
	result, _ := strconv.ParseInt(finalEnergyStr, 10, 0)
	return int(result)
}

func remove(battery []string, removeCount int) []string {
	if removeCount == 0 {
		return battery
	}
	for idx, energy := range battery {
		if idx == len(battery)-1 {
			break
		}
		energyNext := battery[idx+1]
		if energy < energyNext {
			updatedBattery := append(battery[:idx], battery[idx+1:]...)
			return remove(updatedBattery, removeCount-1)
		}
	}
	return remove(battery[:len(battery)-1], removeCount-1)
}
