package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(24)
	expected := 4
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.SampleNr(24, 2)
	expected := 2024
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(24)
	expected := 48806532300520
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(24)
	expected := "ddn,kqh,nhs,nnf,wrc,z09,z20,z34"
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
