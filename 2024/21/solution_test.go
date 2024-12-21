package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(21)
	expected := 126384
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(21)
	expected := 128962
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(21)
	expected := 159684145150108
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
