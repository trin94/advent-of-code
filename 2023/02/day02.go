package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Reveal struct {
	red   int32
	green int32
	blue  int32
}

type Game struct {
	id      int32
	reveals []Reveal
}

func main() {
	file := "2023/02/input.txt"
	lines := readLinesFrom(file)
	games := parseGamesFrom(lines)

	part1Solution := solvePart1(games, 12, 13, 14)
	fmt.Printf("Part 1: %d\n", part1Solution)

	part2Solution := solvePart2(games)
	fmt.Printf("Part 2: %d\n", part2Solution)
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func parseGamesFrom(lines []string) []Game {
	var games []Game

	for _, line := range lines {
		split := strings.Split(line, ":")

		gameId := func() int32 {
			gameIdString := strings.Split(split[0], " ")[1]
			gameId, _ := strconv.ParseInt(gameIdString, 10, 0)
			return int32(gameId)
		}()

		reveals := func() []Reveal {
			var reveals []Reveal
			revealIterations := strings.Split(split[1], ";")

			for _, revealIteration := range revealIterations {
				amountToColorMappings := strings.Split(revealIteration, ",")

				var red int32
				var green int32
				var blue int32

				for _, amountToColorMapping := range amountToColorMappings {
					mapping := strings.Split(strings.TrimSpace(amountToColorMapping), " ")
					amount, _ := strconv.ParseInt(mapping[0], 10, 0)
					color := mapping[1]
					if color == "red" {
						red = int32(amount)
					} else if color == "green" {
						green = int32(amount)
					} else if color == "blue" {
						blue = int32(amount)
					} else {
						panic("Cannot handle color: " + color)
					}
				}
				reveals = append(reveals, Reveal{red: red, green: green, blue: blue})
			}
			return reveals
		}()

		games = append(games, Game{gameId, reveals})
	}

	return games
}

func solvePart1(games []Game, maxRed int32, maxGreen int32, maxBlue int32) int32 {
	var gameIdSum int32

	for _, game := range games {

		allRevealsPossible := func() bool {
			for _, reveal := range game.reveals {
				if reveal.red > maxRed || reveal.green > maxGreen || reveal.blue > maxBlue {
					return false
				}
			}
			return true
		}()

		if allRevealsPossible {
			gameIdSum += game.id
		}
	}

	return gameIdSum
}

func solvePart2(games []Game) int32 {
	var powers int32

	for _, game := range games {
		var atLeastRed int32
		var atLeastGreen int32
		var atLeastBlue int32

		for _, reveal := range game.reveals {
			atLeastRed = max(atLeastRed, reveal.red)
			atLeastGreen = max(atLeastGreen, reveal.green)
			atLeastBlue = max(atLeastBlue, reveal.blue)
		}

		powers += atLeastRed * atLeastGreen * atLeastBlue
	}

	return powers
}
