package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(23)
	expected := 7
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(23)
	expected := 1194
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.Sample(23)
	expected := "co,de,ka,ta"
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(23)
	expected := "bd,bu,dv,gl,qc,rn,so,tm,wf,yl,ys,ze,zr"
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
