package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Sample(t *testing.T) {
	input := "sample.txt"
	expected := 55312
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart1Puzzle(t *testing.T) {
	input := "input.txt"
	expected := 239714
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart2Puzzle(t *testing.T) {
	input := "input.txt"
	expected := 284973560658514
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
