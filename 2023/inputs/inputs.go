package inputs

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const YEAR = 2023
const DIRECTORY = "inputs"

// Sample returns the absolute path to the sample file for the given day
func Sample(day int) string {
	return SampleNr(day, 1)
}

// SampleNr returns the absolute path to the specific sample file for the given day
func SampleNr(day int, sample int) string {
	return buildPath(DIRECTORY, fmt.Sprintf("%02d-sample.%d.txt", day, sample))
}

// Input returns the absolute path to the input file for the given day
func Input(day int) string {
	return InputNr(day, 1)
}

// InputNr returns the absolute path to the specific input file for the given day
func InputNr(day int, sample int) string {
	return buildPath(DIRECTORY, fmt.Sprintf("%02d-input.%d.txt", day, sample))
}

func buildPath(parts ...string) string {
	paths := make([]string, 0, len(parts)+1)
	paths = append(paths, getRootDirectory())
	paths = append(paths, parts...)
	return filepath.Join(paths...)
}

func getRootDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		panic("Cannot get current directory")
	}
	sep := string(os.PathSeparator)
	year := strconv.Itoa(YEAR)
	list := strings.Split(dir, sep)
	var directoryIndex int
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == year {
			directoryIndex = i
			break
		}
	}
	return sep + filepath.Join(list[:directoryIndex+1]...)
}

func ReadLinesFrom(path string) []string {
	inputByteStream, err := os.ReadFile(path)
	if err != nil {
		panic("Could not read file from " + path)
	}
	content := string(inputByteStream)
	content = strings.TrimSpace(content)
	return strings.Split(content, "\n")
}
