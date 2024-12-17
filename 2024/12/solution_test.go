package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestSample1(t *testing.T) {
	input := inputs.SampleNr(12, 1)
	expected := 140
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.SampleNr(12, 2)
	expected := 772
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}
func TestSample3(t *testing.T) {
	input := inputs.SampleNr(12, 3)
	expected := 1930
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(12)
	expected := 1494342
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample4(t *testing.T) {
	input := inputs.SampleNr(12, 1)
	expected := 80
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPart2Puzzle(t *testing.T) {
	input := inputs.Input(12)
	expected := 893676
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
