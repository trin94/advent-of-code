package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(18)
	expected := 22
	actual := solvePuzzle1(input, 6, 12)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(18)
	expected := 268
	actual := solvePuzzle1(input, 70, 1024)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(18)
	expected := "6,1"
	actual := solvePuzzle2(input, 6)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(18)
	expected := "64,11"
	actual := solvePuzzle2(input, 70)
	assert.Equal(t, expected, actual)
}
