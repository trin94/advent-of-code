package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestPart1Sample1(t *testing.T) {
	input := inputs.SampleNr(10, 1)
	expected := 2
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart1Sample2(t *testing.T) {
	input := inputs.SampleNr(10, 2)
	expected := 4
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart1Sample3(t *testing.T) {
	input := inputs.SampleNr(10, 3)
	expected := 3
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart1Sample4(t *testing.T) {
	input := inputs.SampleNr(10, 4)
	expected := 36
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(10)
	expected := 694
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart2Sample(t *testing.T) {
	input := inputs.SampleNr(10, 4)
	expected := 81
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPart2Puzzle2(t *testing.T) {
	input := inputs.Input(10)
	expected := 1497
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
