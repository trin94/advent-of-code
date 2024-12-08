package p10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSamples1(t *testing.T) {
	assert.Equal(t, 4, solvePuzzle1("10.sample.01.txt"))
	assert.Equal(t, 4, solvePuzzle1("10.sample.02.txt"))
	assert.Equal(t, 8, solvePuzzle1("10.sample.03.txt"))
	assert.Equal(t, 8, solvePuzzle1("10.sample.04.txt"))
}

func TestPuzzle1(t *testing.T) {
	input := "10.input.txt"
	expected := 6886
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	assert.Equal(t, 4, solvePuzzle2("10.sample.05.txt"))
	assert.Equal(t, 4, solvePuzzle2("10.sample.06.txt"))
}

func TestPuzzle2(t *testing.T) {
	input := "10.input.txt"
	expected := 371
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
