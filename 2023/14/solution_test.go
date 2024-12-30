package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2023/inputs"
	"trin94/aoc/2023/utils"
)

var testMoveTowardsNorthColumnsParams = []struct {
	input    string
	expected string
}{
	{
		input:    "OO.O.O..##",
		expected: "OOOO....##",
	},
	{
		input:    "...OO....O",
		expected: "OOO.......",
	},
	{
		input:    ".O...#O..O",
		expected: "O....#OO..",
	},
	{
		input:    ".O.#.OO.O.#.O",
		expected: "O..#OOO...#O.",
	},
}

func TestMoveTowardsNorthColumns(t *testing.T) {
	for _, tt := range testMoveTowardsNorthColumnsParams {
		input := utils.RotateClockwise([]string{tt.input})
		expected := utils.RotateClockwise([]string{tt.expected})
		grid := utils.NewGrid(input)
		moveStonesTowardsNorth(&grid)
		assert.Equal(t, expected, grid.Lines())
	}
}

func TestSample1(t *testing.T) {
	input := inputs.Sample(14)
	expected := 136
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(14)
	expected := 108857
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(14)
	expected := 95273
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
