package main

import (
	"testing"
	"trin94/aoc/2023/inputs"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(18)
	expected := 62
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(18)
	expected := 49061
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestUpdatedInstruction(t *testing.T) {
	var params = []struct {
		input    Instruction
		expected Instruction
	}{
		{
			input: Instruction{
				direction: 0,
				amount:    0,
				color:     "#70c710",
			}, expected: Instruction{
				direction: RIGHT,
				amount:    461937,
				color:     "#70c710",
			},
		},
		{
			input: Instruction{
				direction: 0,
				amount:    0,
				color:     "#caa173",
			}, expected: Instruction{
				direction: UP,
				amount:    829975,
				color:     "#caa173",
			},
		},
		{
			input: Instruction{
				direction: 0,
				amount:    0,
				color:     "#015232",
			}, expected: Instruction{
				direction: LEFT,
				amount:    5411,
				color:     "#015232",
			},
		},
		{
			input: Instruction{
				direction: 0,
				amount:    0,
				color:     "#caa171",
			}, expected: Instruction{
				direction: DOWN,
				amount:    829975,
				color:     "#caa171",
			},
		},
	}

	for _, test := range params {
		assert.Equal(t, test.expected, test.input.UpdatedInstruction())
	}
}

// func TestSample2(t *testing.T) {
// 	input := inputs.Sample(18)
// 	expected := 1
// 	actual := solvePuzzle2(input)
// 	assert.Equal(t, expected, actual)
// }

// func TestPuzzle2(t *testing.T) {
// 	input := inputs.Input(18)
// 	expected := 1
// 	actual := solvePuzzle2(input)
// 	assert.Equal(t, expected, actual)
// }
