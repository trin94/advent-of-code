package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

// ###

type Coordinate struct {
	col, row int
}

// ###

type Set[E comparable] map[E]struct{}

func NewSet[E comparable]() Set[E] {
	return Set[E]{}
}

func (s Set[E]) Add(v E) {
	s[v] = struct{}{}
}

func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}

// ###

type Grid struct {
	lines   []string
	columns int
	rows    int
}

func NewGrid(lines []string) Grid {
	return Grid{
		lines:   lines,
		columns: len(lines[0]),
		rows:    len(lines),
	}
}

func (g Grid) charAt(row int, col int) string {
	if g.contains(row, col) {
		return string(g.lines[row][col])
	} else {
		return ""
	}
}

func (g Grid) contains(row int, col int) bool {
	return row >= 0 && row < g.rows && col >= 0 && col < g.columns
}

func (g Grid) locationsOf(frequency string) []Coordinate {
	result := make([]Coordinate, 0)
	for r := 0; r < g.rows; r++ {
		for c := 0; c < g.columns; c++ {
			if g.charAt(r, c) == frequency {
				result = append(result, Coordinate{c, r})
			}
		}
	}
	return result
}

func solvePuzzle1(path string) int {
	lines := readLinesFrom(path)
	grid := NewGrid(lines)
	antiNodes := NewSet[Coordinate]()

	for _, frequency := range frequencyIdentifiers() {
		locations := grid.locationsOf(frequency)

		for i := 0; i < len(locations); i++ {
			for j := 0; j < len(locations); j++ {
				if i == j {
					continue
				}

				first := locations[i]
				second := locations[j]

				rowDirection := 1
				if first.row <= second.row {
					rowDirection = -1
				}

				colDirection := 1
				if first.col <= second.col {
					colDirection = -1
				}

				rowDistance := absoluteValue(first.row - second.row)
				colDistance := absoluteValue(first.col - second.col)

				possibleAntiNodeRow := first.row + rowDistance*rowDirection
				possibleAntiNodeCol := first.col + colDistance*colDirection

				if grid.contains(possibleAntiNodeRow, possibleAntiNodeCol) {
					antiNodes.Add(Coordinate{possibleAntiNodeCol, possibleAntiNodeRow})
				}

			}
		}
	}

	return len(antiNodes)
}

func solvePuzzle2(path string) int {
	lines := readLinesFrom(path)
	grid := NewGrid(lines)
	antiNodes := NewSet[Coordinate]()

	for _, frequency := range frequencyIdentifiers() {
		locations := grid.locationsOf(frequency)

		for i := 0; i < len(locations); i++ {
			for j := 0; j < len(locations); j++ {
				if i == j {
					continue
				}

				first := locations[i]
				second := locations[j]

				rowDirection := 1
				if first.row <= second.row {
					rowDirection = -1
				}

				colDirection := 1
				if first.col <= second.col {
					colDirection = -1
				}

				rowDistance := absoluteValue(first.row - second.row)
				colDistance := absoluteValue(first.col - second.col)

				k := 0

				for {
					possibleAntiNodeRow := first.row + rowDistance*rowDirection*k
					possibleAntiNodeCol := first.col + colDistance*colDirection*k

					if !grid.contains(possibleAntiNodeRow, possibleAntiNodeCol) {
						break
					}

					antiNodes.Add(Coordinate{possibleAntiNodeCol, possibleAntiNodeRow})

					k++
				}
			}
		}
	}

	return len(antiNodes)
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func frequencyIdentifiers() []string {
	size := 26*2 + 10
	result := make([]string, size)
	for i := 0; i < 10; i++ {
		result[i] = fmt.Sprintf("%d", i)
	}
	for letter, i := 'a', 10; letter <= 'z'; letter, i = letter+1, i+1 {
		lowerCase := fmt.Sprintf("%c", letter)
		upperCase := fmt.Sprintf("%c", unicode.ToUpper(letter))
		result[i] = lowerCase
		result[i+26] = upperCase
	}
	return result
}

func absoluteValue(value int) int {
	return int(math.Abs(float64(value)))
}
