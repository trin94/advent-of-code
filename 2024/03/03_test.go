package p03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := "03.sample1.txt"

	expected := int64(161)
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "03.input.txt"

	expected := int64(187194524)
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "03.sample2.txt"

	expected := int64(48)
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := "03.input.txt"

	expected := int64(127092535)
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
