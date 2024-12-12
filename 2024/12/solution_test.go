package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSample1(t *testing.T) {
	input := "sample.1.txt"
	expected := 140
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "sample.2.txt"
	expected := 772
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}
func TestSample3(t *testing.T) {
	input := "sample.3.txt"
	expected := 1930
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "input.txt"
	expected := 1494342
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample4(t *testing.T) {
	input := "sample.1.txt"
	expected := 80
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPart2Puzzle(t *testing.T) {
	input := "input.txt"
	expected := 893676
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
