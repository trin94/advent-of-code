package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2023/inputs"
)

func TestSample1(t *testing.T) {

	fmt.Println("1"[1:])
	input := inputs.Sample(12)
	expected := 21
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(12)
	expected := 7402 // 2**17
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(12)
	expected := 525152
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(12)
	expected := 3384337640277
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
