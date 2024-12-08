package p07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := "07.sample.txt"
	expected := 3749
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "07.input.txt"
	expected := 66343330034722
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {

	input := "07.sample.txt"
	expected := 11387
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := "07.input.txt"
	expected := 637696070419031
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
