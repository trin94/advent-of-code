package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var CharacterOrder = "AKQJT98765432"

type CamelCard struct {
	hand          string
	handFrequency map[rune]int
	bid           int
}

func compareTo() func(a CamelCard, b CamelCard) int {
	return func(a, b CamelCard) int {
		if a.isFiveOfAKind() && b.isFiveOfAKind() {
			return a.compareBySingleCards(b)
		} else if a.isFiveOfAKind() {
			return -1
		} else if b.isFiveOfAKind() {
			return 1
		} else if a.isFourOfAKind() && b.isFourOfAKind() {
			return a.compareBySingleCards(b)
		} else if a.isFourOfAKind() {
			return -1
		} else if b.isFourOfAKind() {
			return 1
		} else if a.isFullHouse() && b.isFullHouse() {
			return a.compareBySingleCards(b)
		} else if a.isFullHouse() {
			return -1
		} else if b.isFullHouse() {
			return 1
		} else if a.isThreeOfAKind() && b.isThreeOfAKind() {
			return a.compareBySingleCards(b)
		} else if a.isThreeOfAKind() {
			return -1
		} else if b.isThreeOfAKind() {
			return 1
		} else if a.isTwoPair() && b.isTwoPair() {
			return a.compareBySingleCards(b)
		} else if a.isTwoPair() {
			return -1
		} else if b.isTwoPair() {
			return 1
		} else if a.isOnePair() && b.isOnePair() {
			return a.compareBySingleCards(b)
		} else if a.isOnePair() {
			return -1
		} else if b.isOnePair() {
			return 1
		} else if a.isHighCard() && b.isHighCard() {
			return a.compareBySingleCards(b)
		} else if a.isHighCard() {
			return -1
		} else if b.isHighCard() {
			return 1
		}
		return 0

	}
}

func (camelCard CamelCard) isFiveOfAKind() bool {
	for _, count := range camelCard.handFrequency {
		if count == 5 {
			return true
		}
	}
	return false
}

func (camelCard CamelCard) isFourOfAKind() bool {
	for _, count := range camelCard.handFrequency {
		if count == 4 {
			return true
		}
	}
	return false
}

func (camelCard CamelCard) isFullHouse() bool {
	var twoShared = false
	var threeShared = false
	for _, count := range camelCard.handFrequency {
		if count == 2 {
			twoShared = true
		}
		if count == 3 {
			threeShared = true
		}
	}
	return twoShared && threeShared
}

func (camelCard CamelCard) isThreeOfAKind() bool {
	var single = false
	var threeShared = false
	for _, count := range camelCard.handFrequency {
		if count == 1 {
			single = true
		}
		if count == 3 {
			threeShared = true
		}
	}
	return single && threeShared
}

func (camelCard CamelCard) isTwoPair() bool {
	var firstPair = ""
	var secondPair = ""
	for character, count := range camelCard.handFrequency {
		if count == 2 {
			if firstPair == "" {
				firstPair = string(character)
			}
			if secondPair == "" && firstPair != string(character) {
				secondPair = string(character)
			}
		}
	}
	return firstPair != "" && secondPair != ""
}

func (camelCard CamelCard) isOnePair() bool {
	var pairFound = false
	var different = 0
	for _, count := range camelCard.handFrequency {
		if count == 2 {
			pairFound = true
		}
		if count == 1 {
			different++
		}
	}
	return pairFound && different == 3
}

func (camelCard CamelCard) isHighCard() bool {
	return len(camelCard.handFrequency) == 5
}

func (camelCard CamelCard) compareBySingleCards(otherCard CamelCard) int {
	for i, character := range camelCard.hand {
		thisCharacter := uint8(character)
		thatCharacter := otherCard.hand[i]
		if thisCharacter == thatCharacter {
			continue
		}

		idxThis := strings.IndexRune(CharacterOrder, character)
		idxThat := strings.Index(CharacterOrder, string(thatCharacter))
		if idxThis < idxThat {
			return -1
		} else {
			return 1
		}
	}
	return 0
}

func main() {
	file := "2023/07/input.txt"
	lines := readLinesFrom(file)
	camelCards := parseCamelCards(lines)

	part1Solution := solvePart1(camelCards)
	fmt.Printf("Part 1: %d\n", part1Solution)

	part2Solution := solvePart2(camelCards)
	fmt.Printf("Part 2: %d\n", part2Solution)
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func parseCamelCards(lines []string) (camelCards []CamelCard) {
	for _, line := range lines {
		fields := strings.Fields(line)
		hand := fields[0]
		bid, _ := strconv.ParseInt(fields[1], 10, 0)
		camelCard := CamelCard{
			hand:          hand,
			handFrequency: countCharacters(hand),
			bid:           int(bid),
		}
		camelCards = append(camelCards, camelCard)
	}
	return
}

func countCharacters(str string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range str {
		frequency[char] = frequency[char] + 1
	}
	return frequency
}

func solvePart1(camelCards []CamelCard) (bids int) {
	slices.SortFunc(camelCards, compareTo())
	slices.Reverse(camelCards)
	for rank, card := range camelCards {
		bids += (rank + 1) * card.bid
	}
	return
}

func solvePart2(camelCards []CamelCard) int {
	CharacterOrder = "AKQT98765432J"
	upgradedCards := make([]CamelCard, len(camelCards))
	for i, card := range camelCards {
		upgradedCards[i] = card.upgrade()
	}
	return solvePart1(upgradedCards)
}

func (camelCard CamelCard) upgrade() CamelCard {
	var frequencyCopy = make(map[rune]int)
	var maxValue int
	var maxValueKey rune
	for k, v := range camelCard.handFrequency {
		frequencyCopy[k] = v
		if k != 'J' && v > maxValue {
			maxValue = v
			maxValueKey = k
		}
	}
	jokerValue, jokerExists := camelCard.handFrequency['J']
	if jokerExists {
		delete(frequencyCopy, 'J')
		frequencyCopy[maxValueKey] = frequencyCopy[maxValueKey] + jokerValue
	}
	return CamelCard{
		hand:          camelCard.hand,
		handFrequency: frequencyCopy,
		bid:           camelCard.bid,
	}
}
