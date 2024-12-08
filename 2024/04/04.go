package p04

import (
	"os"
	"strings"
)

type Grid struct {
	lines   []string
	columns int
	rows    int
}

func newGrid(lines []string) Grid {
	return Grid{
		lines:   lines,
		columns: len(lines[0]),
		rows:    len(lines),
	}
}

func (g Grid) charAt(row int, col int) string {
	if row >= g.rows || row < 0 {
		return ""
	}
	if col >= g.columns || col < 0 {
		return ""
	}
	return string(g.lines[row][col])
}

func solvePuzzle1(path string) (distances int64) {
	lines := readLinesFrom(path)

	grid := newGrid(lines)

	var possibleResults []string

	for r := 0; r < grid.rows; r++ {
		for c := 0; c < grid.columns; c++ {
			if grid.charAt(r, c) == "X" {
				west := "X" + grid.charAt(r, c-1) + grid.charAt(r, c-2) + grid.charAt(r, c-3)
				northWest := "X" + grid.charAt(r-1, c-1) + grid.charAt(r-2, c-2) + grid.charAt(r-3, c-3)
				north := "X" + grid.charAt(r-1, c) + grid.charAt(r-2, c) + grid.charAt(r-3, c)
				northEast := "X" + grid.charAt(r-1, c+1) + grid.charAt(r-2, c+2) + grid.charAt(r-3, c+3)
				east := "X" + grid.charAt(r, c+1) + grid.charAt(r, c+2) + grid.charAt(r, c+3)
				southEast := "X" + grid.charAt(r+1, c+1) + grid.charAt(r+2, c+2) + grid.charAt(r+3, c+3)
				south := "X" + grid.charAt(r+1, c) + grid.charAt(r+2, c) + grid.charAt(r+3, c)
				southWest := "X" + grid.charAt(r+1, c-1) + grid.charAt(r+2, c-2) + grid.charAt(r+3, c-3)

				possibleResults = append(possibleResults, west, northWest, north, northEast,
					east, southEast, south, southWest)
			}
		}
	}

	occurrences := 0

	for _, possibleResult := range possibleResults {
		if possibleResult == "XMAS" {
			occurrences++
		}
	}

	return int64(occurrences)
}

func solvePuzzle2(path string) (score int64) {
	lines := readLinesFrom(path)

	grid := newGrid(lines)

	counter := 0

	for r := 0; r < grid.rows; r++ {
		for c := 0; c < grid.columns; c++ {

			if grid.charAt(r, c) == "A" {

				northWest := grid.charAt(r-1, c-1)
				southWest := grid.charAt(r+1, c-1)
				northEast := grid.charAt(r-1, c+1)
				southEast := grid.charAt(r+1, c+1)

				if northWest == southEast || southWest == northEast {
					continue
				}

				chars := []string{northEast, southEast, northWest, southWest}

				mCount := count(chars, func(s string) bool {
					return s == "M"
				})

				sCount := count(chars, func(s string) bool {
					return s == "S"
				})

				if mCount == 2 && sCount == 2 {
					counter++
				}

			}

		}
	}

	return int64(counter)
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func count[T any](slice []T, f func(T) bool) int {
	count := 0
	for i := range slice {
		if f(slice[i]) {
			count++
		}
	}
	return count
}
