package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trin94/aoc/2024/inputs"
)

func TestMix(t *testing.T) {
	secretNr, value := 42, 15
	expected := 37
	actual := mix(secretNr, value)
	assert.Equal(t, expected, actual)
}

func TestPrune(t *testing.T) {
	secretNr := 100000000
	expected := 16113920
	actual := prune(secretNr)
	assert.Equal(t, expected, actual)
}

func TestEvolve(t *testing.T) {
	begin := 123
	next := []int{
		15887950,
		16495136,
		527345,
		704524,
		1553684,
		12683156,
		11100544,
		12249484,
		7753432,
		5908254,
	}
	for i, nextSecretNr := range next {
		actual := developSecretNr(begin, i+1)
		expected := nextSecretNr
		assert.Equal(t, expected, actual)
	}
}

func TestSample1(t *testing.T) {
	input := inputs.Sample(22)
	expected := 37327623
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle1(t *testing.T) {
	input := inputs.Input(22)
	expected := 19822877190
	actual := solvePuzzle1(input)
	assert.Equal(t, expected, actual)
}

func TestDeterminePrices(t *testing.T) {
	expected := []int{3, 0, 6, 5, 4, 4, 6, 4, 4, 2}
	actual := determineNextPrices(123, len(expected))
	assert.Equal(t, expected, actual)
}

func TestDeterminePriceChanges(t *testing.T) {
	prices := []int{3, 0, 6, 5, 4, 4, 6, 4, 4, 2}
	actual := determinePriceChanges(prices)
	expected := []int{3, -3, 6, -1, -1, 0, 2, -2, 0, -2}
	assert.Equal(t, expected, actual)
}

func TestSample2(t *testing.T) {
	input := inputs.SampleNr(22, 2)
	expected := 23
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}

func TestPuzzle2(t *testing.T) {
	input := inputs.Input(22)
	expected := 2277
	actual := solvePuzzle2(input)
	assert.Equal(t, expected, actual)
}
