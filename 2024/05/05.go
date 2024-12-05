package p05

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

type Set struct {
	values map[string]bool
}

func NewSet() Set {
	return Set{values: make(map[string]bool)}
}

func (s *Set) Add(value string) {
	s.values[value] = true
}

func (s *Set) Contains(value string) bool {
	return s.values[value]
}

func solvePuzzle1(path string) int {
	rules, updates := readInput(path)
	parsedRules := parseRules(rules)

	sum := 0

	for _, update := range updates {

		nrs := strings.Split(update, ",")

		violationFound := false

		for i, nr := range nrs {
			for j := i + 1; j < len(nrs); j++ {

				requiredSuccessors := parsedRules[nrs[j]]
				if requiredSuccessors.Contains(nr) {
					violationFound = true
					break
				}
			}

			if violationFound {
				break
			}

		}

		if !violationFound {
			val, _ := strconv.Atoi(nrs[len(nrs)/2])
			sum += val
		}
	}

	return sum
}

func solvePuzzle2(path string) int {
	rules, updates := readInput(path)
	parsedRules := parseRules(rules)

	sum := 0

	for _, update := range updates {
		nrs := strings.Split(update, ",")
		modifiedLine := false

	lineStart:
		for i := 0; i < len(nrs); i++ {
			nr := nrs[i]
			for j := i + 1; j < len(nrs); j++ {

				requiredSuccessors := parsedRules[nrs[j]]
				if requiredSuccessors.Contains(nr) {
					nrs = slices.Insert(nrs, 0, nrs[j])
					nrs = slices.Delete(nrs, j+1, j+2)
					modifiedLine = true
					goto lineStart
				}

			}
		}

		if modifiedLine {
			val, _ := strconv.Atoi(nrs[len(nrs)/2])
			sum += val
		}
	}

	return sum
}

func readInput(path string) (rules []string, updates []string) {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)

	substrings := strings.SplitN(inputString, "\n\n", 2)

	rules = strings.Split(strings.TrimSpace(substrings[0]), "\n")
	updates = strings.Split(strings.TrimSpace(substrings[1]), "\n")
	return
}

func parseRules(rules []string) map[string]Set {
	mmm := map[string]Set{}
	for _, rule := range rules {
		ruleSplit := strings.Split(rule, "|")

		first := ruleSplit[0]
		second := ruleSplit[1]

		existingValue, ok := mmm[first] // check for existence
		if ok {
			existingValue.Add(second)
		} else {
			s := NewSet()
			s.Add(second)
			mmm[first] = s
		}

	}
	return mmm
}
