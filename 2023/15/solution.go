package main

import (
	"strconv"
	"strings"
	"trin94/aoc/2023/inputs"
)

type Element struct {
	sequence, operation string
	focalLength         int
}

func solvePuzzle1(path string) (result int) {
	lines := inputs.ReadLinesFrom(path)
	steps := strings.Split(lines[0], ",")
	for _, step := range steps {
		result += runHash(step)
	}
	return
}

func runHash(sequence string) (result int) {
	for _, character := range sequence {
		ascii := int(character)
		result += ascii
		result *= 17
		result %= 256
	}
	return
}

func solvePuzzle2(path string) (result int) {
	lines := inputs.ReadLinesFrom(path)
	steps := strings.Split(lines[0], ",")

	hashMap := store(steps)

	for i, box := range hashMap {
		for j, entry := range box {
			result += (i + 1) * (j + 1) * entry.focalLength
		}
	}

	return
}

func store(steps []string) [][]Element {
	hashmap := make([][]Element, 255)
	for i := range hashmap {
		hashmap[i] = make([]Element, 0)
	}

	for _, step := range steps {
		nextElement := parseElement(step)

		operation, sequence := nextElement.operation, nextElement.sequence
		box := runHash(sequence)

		if operation == "-" {
			for i, existing := range hashmap[box] {
				if existing.sequence == sequence {
					hashmap[box] = removeElement(hashmap[box], i)
					break
				}
			}
		} else if operation == "=" {
			updated := false
			for i, existing := range hashmap[box] {
				if existing.sequence == sequence {
					hashmap[box][i] = nextElement
					updated = true
					break
				}
			}
			if !updated {
				hashmap[box] = append(hashmap[box], nextElement)
			}
		}
	}
	return hashmap
}

func parseElement(value string) (element Element) {
	fields := splitCharsInclusive(value, "=-")
	sequence := fields[0]
	operation := fields[1]
	var focalLength int
	if len(fields) > 2 {
		focalLength, _ = strconv.Atoi(fields[2])
	} else {
		focalLength = -1
	}
	return Element{sequence, operation, focalLength}
}

func splitCharsInclusive(sequence, delimiters string) (out []string) {
	for {
		m := strings.IndexAny(sequence, delimiters)
		if m < 0 {
			break
		}
		out = append(out, sequence[:m], sequence[m:m+1])
		sequence = sequence[m+1:]
	}
	out = append(out, sequence)
	return
}

func removeElement[E comparable](slice []E, index int) []E {
	return append(slice[:index], slice[index+1:]...)
}
