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

func (card Card) countWinningDraws() (sum int) {
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

	part2Solution := solvePart2(cards)
	fmt.Printf("Part 2: %d\n", part2Solution)
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

func solvePart2(cards []Card) (totalCards int) {
	state := make(map[int]int, len(cards))

	for id := range cards {
		state[id+1] = 1
	}

	for cardIndex := 0; cardIndex < len(state); cardIndex++ {
		card := cards[cardIndex]
		cardId := cardIndex + 1
		matchCount := card.countWinningDraws()

		for a := cardIndex + 1; a <= cardIndex+matchCount; a++ {
			state[a+1] = state[a+1] + state[cardId]
		}
	}

	for cardIndex := 0; cardIndex < len(state); cardIndex++ {
		totalCards += state[cardIndex+1]
	}

	return
}
