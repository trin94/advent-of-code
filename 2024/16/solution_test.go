package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestDetermineCosts(t *testing.T) {
	assert.Equal(t, 1, determineCost(North, North))
	assert.Equal(t, 1001, determineCost(North, East))
	assert.Equal(t, 2002, determineCost(North, South))
	assert.Equal(t, 1001, determineCost(North, West))

	assert.Equal(t, 1001, determineCost(West, North))
	assert.Equal(t, 2002, determineCost(West, East))
	assert.Equal(t, 1001, determineCost(West, South))
	assert.Equal(t, 1, determineCost(West, West))
}

func TestSample1(t *testing.T) {
	input := inputs.SampleNr(16, 1)
	expected := 7036
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.SampleNr(16, 2)
	expected := 11048
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(16)
	expected := 94444
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample3(t *testing.T) {
	input := inputs.SampleNr(16, 1)
	expected := 45
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestSample4(t *testing.T) {
	input := inputs.SampleNr(16, 2)
	expected := 64
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(16)
	expected := 502
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
