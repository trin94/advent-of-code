package p09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := "09.sample.txt"
	expected := 114
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "09.input.txt"
	expected := 1637452029
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "09.sample.txt"
	expected := 2
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := "09.input.txt"
	expected := 908
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
