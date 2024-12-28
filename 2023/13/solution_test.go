package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2023/inputs"
)

func TestSample1(t *testing.T) {
	input := inputs.Sample(13)
	expected := 405
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.SampleNr(13, 2)
	expected := 709
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(13)
	expected := 35360
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample3(t *testing.T) {
	input := inputs.Sample(13)
	expected := 400
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestSample4(t *testing.T) {
	input := inputs.SampleNr(13, 2)
	expected := 1400
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestSample5(t *testing.T) {
	input := inputs.SampleNr(13, 3)
	expected := 15
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(13)
	expected := 36755
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
