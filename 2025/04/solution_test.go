package main

import (
	"testing"
	"trin94/aoc/2025/inputs"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(04)
	expected := 13
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(04)
	expected := 1320
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(04)
	expected := 43
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(04)
	expected := 8354
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
