package p07

import (
	"testing"
	"trin94/aoc/2024/inputs"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(7)
	expected := 3749
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(7)
	expected := 66343330034722
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(7)
	expected := 11387
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(7)
	expected := 637696070419031
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
