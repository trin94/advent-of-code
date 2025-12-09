package main

import (
	"testing"
	"trin94/aoc/2025/inputs"
	"trin94/aoc/2025/utils"

	"github.com/stretchr/testify/assert"
)

func TestAreaBetween(t *testing.T) {
	inputA := utils.Coordinate{
		Col: 2,
		Row: 5,
	}
	inputB := utils.Coordinate{
		Col: 11,
		Row: 1,
	}
	expected := 50
	actual := areaBetween(inputA, inputB)
	assert.Equal(t, expected, actual)
}

func TestSample1(t *testing.T) {
	input := inputs.Sample(9)
	expected := 50
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(9)
	expected := 4745816424
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(9)
	expected := 24
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(9)
	expected := 1351617690
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
