package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(19)
	expected := 6
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(19)
	expected := 251
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(19)
	expected := 16
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(19)
	expected := 616957151871345
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
