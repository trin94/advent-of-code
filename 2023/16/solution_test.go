package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2023/inputs"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(16)
	expected := 46
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(16)
	expected := 7623
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(16)
	expected := 51
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(16)
	expected := 8244
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
