package main

import (
	"strings"
	"trin94/aoc/2024/inputs"
)

func solvePuzzle1(path string) int {
	lines := inputs.ReadLinesFrom(path)
	locks, keys, lockSize := parseLocksAndKeys(lines)

	isFit := func(lock, key []int) bool {
		if len(lock) != len(key) {
			return false
		}
		for i := range lock {
			if lock[i]+key[i] >= lockSize {
				return false
			}
		}
		return true
	}

	result := 0

	for lockNr := 0; lockNr < len(locks); lockNr++ {
		lock := locks[lockNr]
		for keyNr := 0; keyNr < len(keys); keyNr++ {
			key := keys[keyNr]
			if isFit(lock, key) {
				result++
			}
		}
	}

	return result
}

func parseLocksAndKeys(lines []string) (locks, keys [][]int, lockSize int) {
	batch, isLock := make([]string, 0), false

	processBatch := func() {
		if isLock {
			batch = batch[1:]
		} else {
			batch = batch[:len(batch)-1]
		}

		height, width := len(batch), len(batch[0])
		sizes := make([]int, width)
		for i := 0; i < width; i++ {
			size := 0
			for j := 0; j < height; j++ {
				if batch[j][i] == '#' {
					size++
				}
			}
			sizes[i] = size
		}

		if isLock {
			locks = append(locks, sizes)
		} else {
			keys = append(keys, sizes)
		}
	}

	for i, line := range lines {
		if line == "" {
			processBatch()
			batch = make([]string, 0)
			continue
		}

		if len(batch) == 0 {
			isLock = strings.Count(line, "#") > 0
		}

		batch = append(batch, line)

		if i == len(lines)-1 {
			processBatch()
			lockSize = len(batch)
		}
	}

	return
}
