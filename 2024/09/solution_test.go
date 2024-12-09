package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFilesystemNotation(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []Space{
		{0, true},
		{-1, false},
		{-1, false},
		{1, true},
		{1, true},
		{1, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{-1, false},
		{2, true},
		{2, true},
		{2, true},
		{2, true},
		{2, true},
	}
	actual := convertToFilesystem(input)
	assert.Equal(t, expected, actual)

	input = []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2}
	expected = []Space{
		{0, true},
		{0, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{1, true},
		{1, true},
		{1, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{2, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{3, true},
		{3, true},
		{3, true},
		{-1, false},
		{4, true},
		{4, true},
		{-1, false},
		{5, true},
		{5, true},
		{5, true},
		{5, true},
		{-1, false},
		{6, true},
		{6, true},
		{6, true},
		{6, true},
		{-1, false},
		{7, true},
		{7, true},
		{7, true},
		{-1, false},
		{8, true},
		{8, true},
		{8, true},
		{8, true},
		{9, true},
		{9, true},
	}
	actual = convertToFilesystem(input)
	assert.Equal(t, expected, actual)
}

func TestCompressFilesystemAlgorithm1_a(t *testing.T) {
	input := []Space{
		{0, true},
		{-1, false},
		{1, true},
	}
	expected := []Space{
		{0, true},
		{1, true},
	}
	actual := compressFilesystemAlgorithm1(input)
	assert.Equal(t, expected, actual)
}

func TestCompressFilesystemAlgorithm1_b(t *testing.T) {
	input := []Space{
		{0, true},
		{0, true},
		{-1, false},
		{1, true},
	}
	expected := []Space{
		{0, true},
		{0, true},
		{1, true},
	}
	actual := compressFilesystemAlgorithm1(input)
	assert.Equal(t, expected, actual)
}

func TestCompressFilesystemAlgorithm1_c(t *testing.T) {
	input := []Space{
		{0, true},
		{-1, false},
		{-1, false},
		{1, true},
	}
	expected := []Space{
		{0, true},
		{1, true},
	}
	actual := compressFilesystemAlgorithm1(input)
	assert.Equal(t, expected, actual)
}

func TestCompressFilesystemAlgorithm1_d(t *testing.T) {
	input := []Space{
		{0, true},
		{-1, false},
		{-1, false},
		{-1, false},
	}
	expected := []Space{
		{0, true},
	}
	actual := compressFilesystemAlgorithm1(input)
	assert.Equal(t, expected, actual)
}

func TestCompressFilesystemAlgorithm1_e(t *testing.T) {
	input := []Space{
		{0, true},
		{-1, false},
		{-1, false},
		{1, true},
		{1, true},
		{1, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{-1, false},
		{2, true},
		{2, true},
		{2, true},
		{2, true},
		{2, true},
	}
	expected := []Space{
		{0, true},
		{2, true},
		{2, true},
		{1, true},
		{1, true},
		{1, true},
		{2, true},
		{2, true},
		{2, true},
	}
	actual := compressFilesystemAlgorithm1(input)
	assert.Equal(t, expected, actual)
}

func TestCompressFilesystemAlgorithm1_f(t *testing.T) {
	input := []Space{
		{0, true},
		{0, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{1, true},
		{1, true},
		{1, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{2, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{3, true},
		{3, true},
		{3, true},
		{-1, false},
		{4, true},
		{4, true},
		{-1, false},
		{5, true},
		{5, true},
		{5, true},
		{5, true},
		{-1, false},
		{6, true},
		{6, true},
		{6, true},
		{6, true},
		{-1, false},
		{7, true},
		{7, true},
		{7, true},
		{-1, false},
		{8, true},
		{8, true},
		{8, true},
		{8, true},
		{9, true},
		{9, true},
	}
	expected := []Space{
		{0, true},
		{0, true},
		{9, true},
		{9, true},
		{8, true},
		{1, true},
		{1, true},
		{1, true},
		{8, true},
		{8, true},
		{8, true},
		{2, true},
		{7, true},
		{7, true},
		{7, true},
		{3, true},
		{3, true},
		{3, true},
		{6, true},
		{4, true},
		{4, true},
		{6, true},
		{5, true},
		{5, true},
		{5, true},
		{5, true},
		{6, true},
		{6, true},
	}
	actual := compressFilesystemAlgorithm1(input)
	assert.Equal(t, expected, actual)
}

func TestCompressFilesystemAlgorithm2(t *testing.T) {
	input := []Space{
		{0, true},
		{0, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{1, true},
		{1, true},
		{1, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{2, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{3, true},
		{3, true},
		{3, true},
		{-1, false},
		{4, true},
		{4, true},
		{-1, false},
		{5, true},
		{5, true},
		{5, true},
		{5, true},
		{-1, false},
		{6, true},
		{6, true},
		{6, true},
		{6, true},
		{-1, false},
		{7, true},
		{7, true},
		{7, true},
		{-1, false},
		{8, true},
		{8, true},
		{8, true},
		{8, true},
		{9, true},
		{9, true},
	}
	expected := []Space{ //00992111777.44.333....5555.6666.....8888..
		{0, true},
		{0, true},
		{9, true},
		{9, true},
		{2, true},
		{1, true},
		{1, true},
		{1, true},
		{7, true},
		{7, true},
		{7, true},
		{-1, false},
		{4, true},
		{4, true},
		{-1, false},
		{3, true},
		{3, true},
		{3, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{-1, false},
		{5, true},
		{5, true},
		{5, true},
		{5, true},
		{-1, false},
		{6, true},
		{6, true},
		{6, true},
		{6, true},
		{-1, false},
		{-1, false},
		{-1, false},
		{-1, false},
		{-1, false},
		{8, true},
		{8, true},
		{8, true},
		{8, true},
		{-1, false},
		{-1, false},
	}
	actual := compressFilesystemAlgorithm2(input)
	assert.Equal(t, expected, actual)
}

func TestSample1(t *testing.T) {
	input := "sample.txt"
	expected := 1928
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := "input.txt"
	expected := 6349606724455
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := "sample.txt"
	expected := 2858
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := "input.txt"
	expected := 6376648986651
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
