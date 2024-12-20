package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestSampleSavings1(t *testing.T) {
	input := inputs.Sample(20)
	_, savings := solvePuzzle1(input)

	assert.Equal(t, 14, savings[2])
	assert.Equal(t, 14, savings[4])
	assert.Equal(t, 2, savings[6])
	assert.Equal(t, 4, savings[8])
	assert.Equal(t, 2, savings[10])
	assert.Equal(t, 3, savings[12])
	assert.Equal(t, 1, savings[20])
	assert.Equal(t, 1, savings[36])
	assert.Equal(t, 1, savings[38])
	assert.Equal(t, 1, savings[40])
	assert.Equal(t, 1, savings[64])

	fmt.Println(savings[12])
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(20)
	expected := 1406

	actual, _ := solvePuzzle1(input)

	assert.Equal(t, expected, actual)

}

func TestSampleSavings2(t *testing.T) {
	input := inputs.Sample(20)
	_, saving := solvePuzzle2(input)

	assert.Equal(t, 3, saving[76])
	assert.Equal(t, 4, saving[74])
	assert.Equal(t, 22, saving[72])
	assert.Equal(t, 12, saving[70])
	assert.Equal(t, 14, saving[68])
	assert.Equal(t, 12, saving[66])
	assert.Equal(t, 19, saving[64])
	assert.Equal(t, 20, saving[62])
	assert.Equal(t, 23, saving[60])
	assert.Equal(t, 25, saving[58])
	assert.Equal(t, 39, saving[56])
	assert.Equal(t, 29, saving[54])
	assert.Equal(t, 31, saving[52])
	assert.Equal(t, 32, saving[50])
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(20)
	expected := 1006101
	actual, _ := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
