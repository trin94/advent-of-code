package main

import (
	"testing"
	"trin94/aoc/2024/inputs"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(8)
	expected := 14
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(8)
	expected := 371
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(8)
	expected := 34
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(8)
	expected := 1229
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
