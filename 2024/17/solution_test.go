package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestTextSample1(t *testing.T) {
	computer := Computer{a: 0, b: 0, c: 9, instructionPointer: 0}
	runProgram(&computer, []int{2, 6})
	assert.Equal(t, 1, computer.b)
}

func TestTextSample2(t *testing.T) {
	computer := Computer{a: 10, b: 0, c: 0, instructionPointer: 0}
	output := runProgram(&computer, []int{5, 0, 5, 1, 5, 4})
	assert.Equal(t, "0,1,2", output)
}

func TestTextSample3(t *testing.T) {
	computer := Computer{a: 2024, b: 0, c: 0, instructionPointer: 0}
	output := runProgram(&computer, []int{0, 1, 5, 4, 3, 0})
	assert.Equal(t, "4,2,5,6,7,7,7,7,3,1,0", output)
	assert.Equal(t, 0, computer.a)
}

func TestTextSample4(t *testing.T) {
	computer := Computer{a: 0, b: 29, c: 0, instructionPointer: 0}
	runProgram(&computer, []int{1, 7})
	assert.Equal(t, 26, computer.b)
}

func TestTextSample5(t *testing.T) {
	computer := Computer{a: 0, b: 2024, c: 43690, instructionPointer: 0}
	runProgram(&computer, []int{4, 0})
	assert.Equal(t, 44354, computer.b)
}

func TestSample1(t *testing.T) {
	input := inputs.SampleNr(17, 1)
	expected := "4,6,3,5,6,3,5,2,1,0"
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(17)
	expected := "7,4,2,5,1,4,6,0,4"
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.SampleNr(17, 2)
	expected := 117440
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(17)
	expected := 164278764924605
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
