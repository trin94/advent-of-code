package p03

import (
	"testing"
	"trin94/aoc/2024/inputs"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(3)
	expected := 161
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(3)
	expected := 187194524
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.SampleNr(3, 2)
	expected := 48
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(3)
	expected := 127092535
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
