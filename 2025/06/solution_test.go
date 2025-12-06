package main

import (
	"testing"
	"trin94/aoc/2025/inputs"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(06)
	expected := 4277556
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(06)
	expected := 5877594983578
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(06)
	expected := 3263827
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(06)
	expected := 11159825706149
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
