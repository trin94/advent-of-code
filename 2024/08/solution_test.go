package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := "sample.txt"
	expected := 14
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "input.txt"
	expected := 371
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "sample.txt"
	expected := 34
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := "input.txt"
	expected := 1229
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
