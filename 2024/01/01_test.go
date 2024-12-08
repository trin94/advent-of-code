package p01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := "01.sample.txt"

	expected := int64(11)
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "01.input.txt"

	expected := int64(2264607)
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "01.sample.txt"

	expected := int64(31)
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := "01.input.txt"

	expected := int64(19457120)
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
