package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(11)
	expected := 374
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(11)
	expected := 9403026
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(11)
	expected := 1030
	actual := solvePuzzle2(input, 10)
	assert.Equal(t, expected, actual)

	input = inputs.Sample(11)
	expected = 8410
	actual = solvePuzzle2(input, 100)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(11)
	expected := 543018317006
	actual := solvePuzzle2(input, 1_000_000)
	assert.Equal(t, expected, actual)
}
