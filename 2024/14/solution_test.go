package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(14)
	expected := 12
	actual := solvePuzzle1(input, 11, 7)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(14)
	expected := 217328832
	actual := solvePuzzle1(input, 101, 103)
	assert.Equal(t, expected, actual)
}

func TestPart2Puzzle(t *testing.T) {
	input := inputs.Input(14)
	expected := 7412
	actual := solvePuzzle2(input, 101, 103)
	assert.Equal(t, expected, actual)
}
