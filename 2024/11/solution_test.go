package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestPart1Sample(t *testing.T) {
	input := inputs.Sample(11)
	expected := 55312
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart1Puzzle(t *testing.T) {
	input := inputs.Input(11)
	expected := 239714
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart2Puzzle(t *testing.T) {
	input := inputs.Input(11)
	expected := 284973560658514
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
