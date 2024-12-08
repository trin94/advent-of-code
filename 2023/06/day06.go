package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := "2023/06/input.txt"
	lines := readLinesFrom(file)

	times, records := parse(lines)

	part1Solution := solvePart1(times, records)
	fmt.Printf("Part 1: %d\n", part1Solution)

	part2Solution := solvePart2(times, records)
	fmt.Printf("Part 2: %d\n", part2Solution)
}

func readLinesFrom(path string) []string {
	inputByteStream, _ := os.ReadFile(path)
	inputString := string(inputByteStream)
	inputString = strings.TrimSpace(inputString)
	return strings.Split(inputString, "\n")
}

func parse(lines []string) (times []int, records []int) {
	for _, str := range strings.Fields(lines[0])[1:] {
		value, _ := strconv.ParseInt(str, 10, 0)
		times = append(times, int(value))
	}
	for _, str := range strings.Fields(lines[1])[1:] {
		value, _ := strconv.ParseInt(str, 10, 0)
		records = append(records, int(value))
	}
	return
}

func solvePart1(times []int, records []int) int {
	var possibleWaysToBeatRecord = 1
	for timesIdx, time := range times {
		record := records[timesIdx]

		beatIt := 0
		for buttonPushedDown := 0; buttonPushedDown <= time; buttonPushedDown++ {
			speed := buttonPushedDown
			rollingTime := time - buttonPushedDown
			if speed*rollingTime > record {
				beatIt++
			}
		}
		possibleWaysToBeatRecord *= beatIt
	}
	return possibleWaysToBeatRecord
}

func solvePart2(times []int, records []int) int {
	t := ""
	for _, time := range times {
		t += strconv.Itoa(time)
	}
	time, _ := strconv.ParseInt(t, 10, 0)
	r := ""
	for _, record := range records {
		r += strconv.Itoa(record)
	}
	record, _ := strconv.ParseInt(r, 10, 0)
	return solvePart1([]int{int(time)}, []int{int(record)})
}
