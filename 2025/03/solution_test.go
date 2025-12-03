package main

import (
	"testing"
	"trin94/aoc/2025/inputs"

	"github.com/stretchr/testify/assert"
)

var testCalculateJoltageParams = []struct {
	battery  string
	expected int
}{
	{
		battery:  "987654321111111",
		expected: 98,
	},
	{
		battery:  "811111111111119",
		expected: 89,
	},
	{
		battery:  "234234234234278",
		expected: 78,
	},
	{
		battery:  "818181911112111",
		expected: 92,
	},
}

func TestCalculateJoltage(t *testing.T) {
	for _, tt := range testCalculateJoltageParams {
		assert.Equal(t, tt.expected, calculateJoltage(tt.battery, 2))
	}
}

func TestSample1(t *testing.T) {
	input := inputs.Sample(03)
	expected := 357
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(03)
	expected := 17144
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

var testCalculateJoltageAdvancedParams = []struct {
	battery  string
	expected int
}{
	{
		battery:  "987654321111111",
		expected: 987654321111,
	},
	{
		battery:  "811111111111119",
		expected: 811111111119,
	},
	{
		battery:  "234234234234278",
		expected: 434234234278,
	},
	{
		battery:  "214234234234278",
		expected: 434234234278,
	},
	{
		battery:  "818181911112111",
		expected: 888911112111,
	},
}

func TestCalculateJoltageAdvanced(t *testing.T) {
	for _, tt := range testCalculateJoltageAdvancedParams {
		assert.Equal(t, tt.expected, calculateJoltage(tt.battery, 12))
	}
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(03)
	expected := 3121910778619
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(03)
	expected := 170371185255900
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
