package p04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := "04.sample.txt"

	expected := int64(18)
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "04.input.txt"

	expected := int64(2397)
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "04.sample.txt"

	expected := int64(9)
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := "04.input.txt"

	expected := int64(1824)
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
