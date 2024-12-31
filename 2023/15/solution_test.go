package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2023/inputs"
)

func TestHash(t *testing.T) {
	input := "HASH"
	expected := 52
	actual := runHash(input)
	assert.Equal(t, expected, actual)
}

func TestSample1(t *testing.T) {
	input := inputs.Sample(15)
	expected := 1320
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(15)
	expected := 519603
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(15)
	expected := 145
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(15)
	expected := 244342
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
