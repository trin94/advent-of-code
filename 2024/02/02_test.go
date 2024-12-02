package p02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := "02.sample.txt"

	expected := int64(2)
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "02.input.txt"

	expected := int64(242)
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "02.sample.txt"

	expected := int64(4)
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := "02.input.txt"

	expected := int64(311)
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
