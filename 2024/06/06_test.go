package p06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := "06.sample.txt"

	expected := 41
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "06.input.txt"

	expected := 4711
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "06.sample.txt"

	expected := 6
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := "06.input.txt"

	expected := 1562
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
