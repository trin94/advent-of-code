package main

import (
	"testing"
	"trin94/aoc/2025/inputs"

	"github.com/stretchr/testify/assert"
)

func TestEuclideanDistance(t *testing.T) {
	first := NewCoordinate(162, 817, 812)
	second := NewCoordinate(425, 690, 689)
	actual := first.EuclideanDistance(second)
	assert.Equal(t, 316.90, actual)
}

func TestSample1(t *testing.T) {
	input := inputs.Sample(8)
	expected := 40
	actual := solvePuzzle1(input, 10)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(8)
	expected := 67488
	actual := solvePuzzle1(input, 1000)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(8)
	expected := 25272
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(8)
	expected := 3767453340
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
