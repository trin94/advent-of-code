package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestParseLocksAndKeys(t *testing.T) {
	input := inputs.Sample(25)
	lines := inputs.ReadLinesFrom(input)
	locks, keys, _ := parseLocksAndKeys(lines)
	expectedLocks := [][]int{
		{0, 5, 3, 4, 3},
		{1, 2, 0, 5, 3},
	}
	expectedKeys := [][]int{
		{5, 0, 2, 1, 3},
		{4, 3, 4, 0, 2},
		{3, 0, 2, 0, 1},
	}
	assert.Equal(t, expectedLocks, locks)
	assert.Equal(t, expectedKeys, keys)
}

func TestSample1(t *testing.T) {
	input := inputs.Sample(25)
	expected := 3
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(25)
	expected := 3466
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}
