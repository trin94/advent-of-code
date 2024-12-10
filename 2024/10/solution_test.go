package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Sample1(t *testing.T) {
	input := "sample.1.txt"
	expected := 2
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart1Sample2(t *testing.T) {
	input := "sample.2.txt"
	expected := 4
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart1Sample3(t *testing.T) {
	input := "sample.3.txt"
	expected := 3
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart1Sample4(t *testing.T) {
	input := "sample.4.txt"
	expected := 36
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "input.txt"
	expected := 694
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPart2Sample(t *testing.T) {
	input := "sample.4.txt"
	expected := 81
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPart2Puzzle2(t *testing.T) {
	input := "input.txt"
	expected := 1497
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
