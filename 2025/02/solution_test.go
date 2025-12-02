package main

import (
	"testing"
	"trin94/aoc/2025/inputs"

	"github.com/stretchr/testify/assert"
)

var testInvalidIdsInRangeParams = []struct {
	line     string
	expected []int
}{
	{
		line:     "11-22",
		expected: []int{11, 22},
	},
	{
		line:     "95-115",
		expected: []int{99},
	},
	{
		line:     "998-1012",
		expected: []int{1010},
	},
	{
		line:     "1188511880-1188511890",
		expected: []int{1188511885},
	},
	{
		line:     "222220-222224",
		expected: []int{222222},
	},
	{
		line:     "1698522-1698528",
		expected: []int{},
	},
	{
		line:     "446443-446449",
		expected: []int{446446},
	},
	{
		line:     "38593856-38593862",
		expected: []int{38593859},
	},
}

func TestInvalidIdsInRange(t *testing.T) {
	for _, tt := range testInvalidIdsInRangeParams {
		r := newRange(tt.line)
		assert.Equal(t, tt.expected, r.invalidIds())
	}
}

func TestSample1(t *testing.T) {
	input := inputs.Sample(02)
	expected := 1227775554
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(02)
	expected := 18952700150
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

var testInvalidIdsExtendedInRangeParams = []struct {
	line     string
	expected []int
}{
	{
		line:     "11-22",
		expected: []int{11, 22},
	},
	{
		line:     "95-115",
		expected: []int{99, 111},
	},
	{
		line:     "998-1012",
		expected: []int{999, 1010},
	},
	{
		line:     "1188511880-1188511890",
		expected: []int{1188511885},
	},
	{
		line:     "222220-222224",
		expected: []int{222222},
	},
	{
		line:     "446443-446449",
		expected: []int{446446},
	},
	{
		line:     "38593856-38593862",
		expected: []int{38593859},
	},
	{
		line:     "565653-565659",
		expected: []int{565656},
	},
	{
		line:     "824824821-824824827",
		expected: []int{824824824},
	},
	{
		line:     "2121212118-2121212124",
		expected: []int{2121212121},
	},
}

func TestInvalidIdsExtendedInRange(t *testing.T) {
	for _, tt := range testInvalidIdsExtendedInRangeParams {
		r := newRange(tt.line)
		assert.Equal(t, tt.expected, r.invalidIdsExtended())
	}
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(02)
	expected := 4174379265
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(02)
	expected := 28858486244
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
