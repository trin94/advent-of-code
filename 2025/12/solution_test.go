package main

import (
	"testing"
	"trin94/aoc/2025/inputs"
	"trin94/aoc/2025/utils"

	"github.com/stretchr/testify/assert"
)

func TestParseShape(t *testing.T) {
	input := []string{
		"0:",
		"###",
		"##.",
		"##.",
	}
	expected := Shape{
		index: 0,
		grid: utils.Grid{
			Columns: 3,
			Rows:    3,
		},
	}
	actual := NewShape(input)
	assert.Equal(t, expected.index, actual.index)
	assert.Equal(t, expected.grid.Columns, actual.grid.Columns)
	assert.Equal(t, expected.grid.Rows, actual.grid.Rows)
}

func TestParseRegion(t *testing.T) {
	var params = []struct {
		input    string
		expected Region
	}{
		{
			input: "12x5: 1 0 1 0 2 2",
			expected: Region{
				width:     12,
				length:    5,
				indexList: []int{1, 0, 1, 0, 2, 2},
			},
		},
		{
			input: "4x4: 0 0 0 0 2 0",
			expected: Region{
				width:     4,
				length:    4,
				indexList: []int{0, 0, 0, 0, 2, 0},
			},
		},
	}

	for _, test := range params {
		assert.Equal(t, test.expected, NewRegion(test.input))
	}
}

func TestPuzzle(t *testing.T) {
	input := inputs.Input(12)
	expected := 479
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}
