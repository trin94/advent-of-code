package main

import (
	"os"
	"strconv"
	"strings"
)

type Space struct {
	value  int
	isFile bool
}

func solvePuzzle1(path string) int {
	numbers := readNumbersFrom(path)
	filesystem := convertToFilesystem(numbers)
	compressedFilesystem := compressFilesystemAlgorithm1(filesystem)

	return calculateChecksum(compressedFilesystem)
}

func solvePuzzle2(path string) int {
	numbers := readNumbersFrom(path)
	filesystem := convertToFilesystem(numbers)
	compressedFilesystem := compressFilesystemAlgorithm2(filesystem)

	return calculateChecksum(compressedFilesystem)

}

func readNumbersFrom(path string) []int {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	numbers := make([]int, len(inputString))
	for i, value := range inputString {
		numbers[i], _ = strconv.Atoi(string(value))
	}
	return numbers
}

func convertToFilesystem(input []int) []Space {
	emptySpace := Space{value: -1, isFile: false}
	output := make([]Space, 0)
	idCounter := 0
	for i, value := range input {
		isFile := i%2 == 0
		for j := 0; j < value; j++ {
			var space Space
			if isFile {
				space = Space{value: idCounter, isFile: true}
			} else {
				space = emptySpace
			}
			output = append(output, space)
		}
		if isFile {
			idCounter++
		}
	}
	return output
}

func compressFilesystemAlgorithm1(input []Space) []Space {
	output := make([]Space, len(input)*10)
	i, j := 0, len(input)-1
	for i <= j {
		iV := input[i]
		if iV.isFile {
			output[i] = iV
			i++
			continue
		}
		jV := input[j]
		for !jV.isFile && i < j {
			j--
			jV = input[j]
		}
		if j <= i {
			break
		}
		output[i] = jV
		i++
		j--
	}
	return output[0:i]
}

func calculateChecksum(compressedFilesystem []Space) int {
	sum := 0
	for i, file := range compressedFilesystem {
		if !file.isFile {
			continue
		}
		sum += file.value * i
	}
	return sum
}

func compressFilesystemAlgorithm2(input []Space) []Space {
	output := clone(input)
	i, j := 0, len(output)-1
	for j > 0 && i < j {

		// scroll to next file
		jV := output[j]
		for !jV.isFile {
			j--
			jV = output[j]
		}

		// determine required space
		upperJ := j
		currentID := jV.value
		lowerJ := j - 1
		for output[lowerJ].isFile && output[lowerJ].value == currentID {
			lowerJ--
			if lowerJ == 0 {
				return output
			}
		}
		spaceRequired := upperJ - lowerJ

		nextI, gapBegin, gapEnd, gapFound := determineNextGap(output, spaceRequired, i, j)
		if gapFound {
			for k, n := gapBegin, 1; k < gapEnd; k, n = k+1, n+1 {
				idxFrom := lowerJ + n
				output[k] = output[idxFrom]
				output[idxFrom] = Space{value: -1, isFile: false}
			}
		}
		j = lowerJ
		i = nextI
	}
	return output
}

func determineNextGap(input []Space, size int, starting int, ending int) (nextSearchBegin, gapBegin, gapEnd int, found bool) {
	nextSearchBegin = starting
	firstGapFound := false
	for i := starting; i < ending; {
		file := input[i]
		if file.isFile {
			i++
			if !firstGapFound {
				nextSearchBegin++
			}
			continue
		}
		firstGapFound = true
		begin := i
		for !file.isFile {
			if i-begin == size {
				return nextSearchBegin, begin, i, true
			}
			i++
			if i >= len(input) {
				return nextSearchBegin, -1, -1, false
			}
			file = input[i]
			if i-begin == size {
				return nextSearchBegin, begin, i, true
			}
		}
	}
	return nextSearchBegin, 1 - 1, -1, false
}

func clone[E comparable](slice []E) []E {
	into := make([]E, len(slice))
	copy(into, slice)
	return into
}
