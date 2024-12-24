package main

import (
	"maps"
	"strconv"
	"trin94/aoc/2024/inputs"
)

type PriceDeltaSequence struct {
	a, b, c, d int
}

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	numbers := parseNumbers(lines)

	sum := 0

	for _, number := range numbers {
		sum += developSecretNr(number, 2000)
	}

	return sum
}

func solvePuzzle2(path string) int {
	lines := inputs.ReadLinesFrom(path)
	numbers := parseNumbers(lines)

	collectedProfits := make(map[PriceDeltaSequence]int)

	for _, number := range numbers {
		prices := determineNextPrices(number, 2000)
		changes := determinePriceChanges(prices)
		profits := determineProfits(prices, changes)

		for sequence, profit := range profits {
			if value, found := collectedProfits[sequence]; found {
				collectedProfits[sequence] = value + profit
			} else {
				collectedProfits[sequence] = profit
			}
		}
	}

	maxProfit := 0
	for profit := range maps.Values(collectedProfits) {
		maxProfit = max(maxProfit, profit)
	}

	return maxProfit
}

func parseNumbers(lines []string) []int {
	result := make([]int, len(lines))
	for i, line := range lines {
		number, _ := strconv.Atoi(line)
		result[i] = number
	}
	return result
}

func developSecretNr(nr, times int) int {
	for range times {
		nr = developSecretNrOnce(nr)
	}
	return nr
}

func developSecretNrOnce(nr int) int {
	nr = prune(mix(nr, nr*64))
	nr = prune(mix(nr, nr/32))
	nr = prune(mix(nr, nr*2048))
	return nr
}

func prune(secretNr int) int {
	return secretNr % 16777216
}

func mix(secretNr, value int) int {
	return secretNr ^ value
}

func determineNextPrices(initial, amount int) []int {
	result := make([]int, amount)
	nr := initial
	result[0] = initial % 10
	for i := 1; i < amount; i++ {
		nr = developSecretNrOnce(nr)
		result[i] = nr % 10
	}
	return result
}

func determinePriceChanges(prices []int) []int {
	result := make([]int, len(prices))
	result[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		result[i] = prices[i] - prices[i-1]
	}
	return result
}

func determineProfits(prices, priceChanges []int) map[PriceDeltaSequence]int {
	sliceSize, result := 4, make(map[PriceDeltaSequence]int)
	for i := sliceSize - 1; i < len(priceChanges); i++ {
		a := priceChanges[i-3]
		b := priceChanges[i-2]
		c := priceChanges[i-1]
		d := priceChanges[i-0]
		key := PriceDeltaSequence{a, b, c, d}
		if _, found := result[key]; !found {
			result[key] = prices[i]
		}
	}
	return result
}
