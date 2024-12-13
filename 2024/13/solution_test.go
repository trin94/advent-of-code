package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSample1(t *testing.T) {
	input := "sample.txt"
	expected := 480
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "input.txt"
	expected := 31761
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "sample.txt"
	expected := 875318608908
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPart2Puzzle(t *testing.T) {
	input := "input.txt"
	expected := 90798500745591
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
