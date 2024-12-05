package p05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := "05.sample.txt"

	expected := 143
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "05.input.txt"

	expected := 4135
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "05.sample.txt"

	expected := 123
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := "05.input.txt"

	expected := 5285
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
