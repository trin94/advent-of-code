package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSample1(t *testing.T) {
	input := "sample.txt"
	expected := 12
	actual := solvePuzzle1(input, 11, 7)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "input.txt"
	expected := 217328832
	actual := solvePuzzle1(input, 101, 103)
	assert.Equal(t, expected, actual)
}

func TestPart2Puzzle(t *testing.T) {
	input := "input.txt"
	expected := 7412
	actual := solvePuzzle2(input, 101, 103)
	assert.Equal(t, expected, actual)
}
