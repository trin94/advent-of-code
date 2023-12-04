package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Card struct {
	winningNumbers []string
	drawnNumbers   []string
}

func (card Card) countWinningDraws() (sum int32) {
	for _, drawn := range card.drawnNumbers {
		for _, winning := range card.winningNumbers {
			if drawn == winning {
				sum += 1
				continue
			}
		}
	}
	return
}

func main() {
	file := "2023/04/input.txt"
	lines := readLinesFrom(file)
	cards := parseCardsFrom(lines)

	part1Solution := solvePart1(cards)
	fmt.Printf("Part 1: %d\n", part1Solution)

	part2Solution := solvePart2()
	fmt.Printf("Part 1: %d\n", part2Solution)
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func parseCardsFrom(lines []string) (cards []Card) {
	for _, line := range lines {
		game := strings.Split(strings.Split(strings.ReplaceAll(line, "  ", " "), ":")[1], "|")

		winning := strings.Split(strings.Trim(game[0], " "), " ")
		drawn := strings.Split(strings.Trim(game[1], " "), " ")

		cards = append(cards, Card{
			winningNumbers: winning,
			drawnNumbers:   drawn,
		})
	}
	return
}

func solvePart1(cards []Card) (result int32) {
	for _, card := range cards {
		countWinningDraws := card.countWinningDraws()
		points := int32(math.Pow(2, float64(countWinningDraws-1)))
		result += points
	}
	return
}

func solvePart2() int32 {
	return 0
}
